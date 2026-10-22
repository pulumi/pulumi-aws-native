// Copyright 2026, Pulumi Corporation.

package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/refdb"
)

const serviceSpecRegistry = "https://registry.npmjs.org/@aws-cdk/aws-service-spec"
const maxServiceSpecBytes = 64 << 20

var serviceSpecVersionPattern = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

type serviceSpecPin struct {
	Version   string `json:"version"`
	Integrity string `json:"integrity"`
}

// CDK's serialized database is an internal format. Only decode the fields we use,
// but validate the table layout so upstream format changes cannot empty our data.
type serviceSpecDatabase struct {
	Schema struct {
		Resource struct {
			Type     string                `json:"type"`
			Entities []serviceSpecResource `json:"entities"`
		} `json:"resource"`
	} `json:"schema"`
}

type serviceSpecResource struct {
	CloudFormationType string   `json:"cloudFormationType"`
	PrimaryIdentifier  []string `json:"primaryIdentifier"`
	// CDK defaults to primaryIdentifier only when cfnRefIdentifier is absent.
	// A pointer preserves an explicit empty list so we do not infer a mapping.
	CfnRefIdentifier *[]string                  `json:"cfnRefIdentifier"`
	Properties       map[string]json.RawMessage `json:"properties"`
	Attributes       map[string]json.RawMessage `json:"attributes"`
}

type refDBOverride struct {
	Reason     string                 `json:"reason"`
	RefReturns metadata.CfRefBehavior `json:"refReturns"`
}

type refDBOverrideReport struct {
	Reason   string                  `json:"reason"`
	Upstream *metadata.CfRefBehavior `json:"upstream"`
	Override metadata.CfRefBehavior  `json:"override"`
}

type refDBReport struct {
	ServiceSpec   serviceSpecPin                 `json:"serviceSpec"`
	ResourceCount int                            `json:"resourceCount"`
	Skipped       map[string]string              `json:"skipped"`
	Overrides     map[string]refDBOverrideReport `json:"overrides"`
}

type refDBChange struct {
	Before *metadata.CfRefBehavior `json:"before"`
	After  *metadata.CfRefBehavior `json:"after"`
}

type refDBOptions struct {
	Update   bool
	Check    bool
	Database string
}

type serviceSpecSource struct {
	client   *http.Client
	registry string
}

func readServiceSpecBytes(reader io.Reader) ([]byte, error) {
	data, err := io.ReadAll(io.LimitReader(reader, maxServiceSpecBytes+1))
	if err != nil {
		return nil, err
	}
	if len(data) > maxServiceSpecBytes {
		return nil, errors.New("service spec exceeds size limit")
	}
	return data, nil
}

func (source serviceSpecSource) fetch(ctx context.Context, url string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := source.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("download service spec: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download %s: %s", url, response.Status)
	}
	return readServiceSpecBytes(response.Body)
}

func (source serviceSpecSource) latest(ctx context.Context) (serviceSpecPin, error) {
	data, err := source.fetch(ctx, source.registry+"/latest")
	if err != nil {
		return serviceSpecPin{}, err
	}
	var release struct {
		Version string `json:"version"`
		Dist    struct {
			Integrity string `json:"integrity"`
		} `json:"dist"`
	}
	if err := json.Unmarshal(data, &release); err != nil {
		return serviceSpecPin{}, fmt.Errorf("decode latest service spec release: %w", err)
	}
	return serviceSpecPin{Version: release.Version, Integrity: release.Dist.Integrity}, nil
}

func (source serviceSpecSource) load(ctx context.Context, pin serviceSpecPin) (serviceSpecDatabase, error) {
	if !serviceSpecVersionPattern.MatchString(pin.Version) {
		return serviceSpecDatabase{}, fmt.Errorf("invalid service spec version %q", pin.Version)
	}
	archive, err := source.fetch(ctx, source.registry+"/-/aws-service-spec-"+pin.Version+".tgz")
	if err != nil {
		return serviceSpecDatabase{}, err
	}
	return decodeServiceSpecArchive(archive, pin.Integrity)
}

func decodeServiceSpecArchive(archive []byte, integrity string) (serviceSpecDatabase, error) {
	digest := sha512.Sum512(archive)
	if "sha512-"+base64.StdEncoding.EncodeToString(digest[:]) != integrity {
		return serviceSpecDatabase{}, errors.New("service spec archive integrity mismatch")
	}
	compressed, err := gzip.NewReader(bytes.NewReader(archive))
	if err != nil {
		return serviceSpecDatabase{}, fmt.Errorf("open service spec archive: %w", err)
	}
	defer compressed.Close()
	entries := tar.NewReader(io.LimitReader(compressed, maxServiceSpecBytes+1))
	for {
		header, err := entries.Next()
		if errors.Is(err, io.EOF) {
			return serviceSpecDatabase{}, errors.New("service spec archive has no package/db.json.gz")
		}
		if err != nil {
			return serviceSpecDatabase{}, fmt.Errorf("read service spec archive: %w", err)
		}
		if header.Name != "package/db.json.gz" {
			continue
		}
		if header.Typeflag != tar.TypeReg {
			return serviceSpecDatabase{}, errors.New("service spec database is not a regular file")
		}
		database, err := gzip.NewReader(entries)
		if err != nil {
			return serviceSpecDatabase{}, fmt.Errorf("open service spec database: %w", err)
		}
		defer database.Close()
		data, err := readServiceSpecBytes(database)
		if err != nil {
			return serviceSpecDatabase{}, fmt.Errorf("read service spec database: %w", err)
		}
		var result serviceSpecDatabase
		if err := json.Unmarshal(data, &result); err != nil {
			return serviceSpecDatabase{}, fmt.Errorf("decode service spec database: %w", err)
		}
		return result, nil
	}
}

func generateRefDB(
	database serviceSpecDatabase, overrides map[string]refDBOverride,
) (*refdb.RefDB, refDBReport, error) {
	table := database.Schema.Resource
	if table.Type != "entities" || len(table.Entities) == 0 {
		return nil, refDBReport{}, errors.New("unsupported or empty CDK resource table")
	}
	result := &refdb.RefDB{Resources: map[string]refdb.RefDBResource{}}
	report := refDBReport{Skipped: map[string]string{}, Overrides: map[string]refDBOverrideReport{}}
	seen := map[string]bool{}
	for _, resource := range table.Entities {
		name := resource.CloudFormationType
		if name == "" || seen[name] {
			return nil, refDBReport{}, fmt.Errorf("missing or duplicate resource name %q", name)
		}
		seen[name] = true
		identifiers := resource.PrimaryIdentifier
		if resource.CfnRefIdentifier != nil {
			identifiers = *resource.CfnRefIdentifier
		}
		if len(identifiers) == 0 {
			report.Skipped[name] = "No Ref or primary identifier"
			continue
		}
		if resource.Properties == nil || resource.Attributes == nil {
			return nil, refDBReport{}, fmt.Errorf("missing properties or attributes in CDK resource %s", name)
		}
		var missing []string
		for _, property := range identifiers {
			if property == "" {
				return nil, refDBReport{}, fmt.Errorf("invalid identifier for %s", name)
			}
			_, isProperty := resource.Properties[property]
			_, isAttribute := resource.Attributes[property]
			if !isProperty && !isAttribute {
				missing = append(missing, property)
			}
		}
		if len(missing) > 0 {
			// CfRefBehavior only supports top-level fields, not nested paths.
			report.Skipped[name] = "Identifiers unavailable as top-level properties: " + strings.Join(missing, ", ")
			continue
		}
		behavior := metadata.CfRefBehavior{Property: identifiers[0]}
		if len(identifiers) > 1 {
			behavior = metadata.CfRefBehavior{Properties: identifiers, Delimiter: "|"}
		}
		result.Resources[name] = refdb.RefDBResource{RefReturns: refdb.RefDBRecord{CfRefBehavior: behavior}}
	}
	for name, override := range overrides {
		if strings.TrimSpace(override.Reason) == "" {
			return nil, refDBReport{}, fmt.Errorf("override for %s needs an evidence-backed reason", name)
		}
		behavior := override.RefReturns
		kinds := 0
		for _, set := range []bool{
			behavior.Property != "", len(behavior.Properties) > 0, behavior.NotSupported, behavior.NotSupportedYet,
		} {
			if set {
				kinds++
			}
		}
		if kinds != 1 || (len(behavior.Properties) > 0 && behavior.Delimiter == "") {
			return nil, refDBReport{}, fmt.Errorf("invalid Ref override for %s", name)
		}
		var upstream *metadata.CfRefBehavior
		if resource, ok := result.Resources[name]; ok {
			upstream = &resource.RefReturns.CfRefBehavior
		}
		report.Overrides[name] = refDBOverrideReport{Reason: override.Reason, Upstream: upstream, Override: behavior}
		result.Resources[name] = refdb.RefDBResource{RefReturns: refdb.RefDBRecord{CfRefBehavior: behavior}}
		delete(report.Skipped, name)
	}
	if len(result.Resources) == 0 {
		return nil, refDBReport{}, errors.New("service spec yielded no Ref mappings")
	}
	report.ResourceCount = len(result.Resources)
	return result, report, nil
}

func compareRefDB(previous, current *refdb.RefDB) map[string]refDBChange {
	changes := map[string]refDBChange{}
	for name, resource := range previous.Resources {
		before := resource.RefReturns.CfRefBehavior
		if next, ok := current.Resources[name]; ok {
			after := next.RefReturns.CfRefBehavior
			if !reflect.DeepEqual(before, after) {
				changes[name] = refDBChange{Before: &before, After: &after}
			}
		} else {
			changes[name] = refDBChange{Before: &before}
		}
	}
	for name, resource := range current.Resources {
		if _, ok := previous.Resources[name]; !ok {
			after := resource.RefReturns.CfRefBehavior
			changes[name] = refDBChange{After: &after}
		}
	}
	return changes
}

func encodeRefDBJSON(value any) ([]byte, error) {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(data, '\n'), nil
}

func readRefDBJSON(path string, value any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, value); err != nil {
		return fmt.Errorf("decode %s: %w", path, err)
	}
	return nil
}

func runRefDB(ctx context.Context, root string, options refDBOptions, source serviceSpecSource) error {
	if options.Update && (options.Check || options.Database != "") {
		return errors.New("--update cannot be combined with --check or --database")
	}
	pinPath := filepath.Join(root, "meta", "service-spec-version.json")
	var pin serviceSpecPin
	if err := readRefDBJSON(pinPath, &pin); err != nil {
		return err
	}
	if options.Update {
		latest, err := source.latest(ctx)
		if err != nil {
			return err
		}
		pin = latest
	}
	var database serviceSpecDatabase
	if options.Database != "" {
		if err := readRefDBJSON(options.Database, &database); err != nil {
			return err
		}
	} else {
		loaded, err := source.load(ctx, pin)
		if err != nil {
			return err
		}
		database = loaded
	}
	var overrides struct {
		Resources map[string]refDBOverride `json:"resources"`
	}
	overridePath := filepath.Join(root, "meta", "ref-db-overrides.json")
	overrideData, err := os.ReadFile(overridePath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(bytes.NewReader(overrideData))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&overrides); err != nil {
		return fmt.Errorf("decode %s: %w", overridePath, err)
	}
	result, report, err := generateRefDB(database, overrides.Resources)
	if err != nil {
		return err
	}
	report.ServiceSpec = pin
	target := filepath.Join(root, "meta", "ref-db.json")
	previous := new(refdb.RefDB)
	if err := previous.LoadJSONFile(target); err != nil {
		return err
	}
	changes := compareRefDB(previous, result)
	reportData, err := encodeRefDBJSON(report)
	if err != nil {
		return err
	}
	reportPath := filepath.Join(root, "reports", "ref-db.json")
	if options.Check {
		existing, err := os.ReadFile(reportPath)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
		if !reflect.DeepEqual(previous, result) || !bytes.Equal(existing, reportData) {
			return errors.New("ref database or report is stale. Run mise exec -- make ref-db")
		}
	} else {
		type output struct {
			path  string
			value any
		}
		outputs := []output{
			{target, result},
			{reportPath, report},
			{filepath.Join(root, ".runtime", "ref-db-changes.json"), changes},
		}
		if options.Update {
			// Save the new pin only after its derived outputs have been written.
			outputs = append(outputs, output{pinPath, pin})
		}
		for _, output := range outputs {
			data, err := encodeRefDBJSON(output.value)
			if err != nil {
				return err
			}
			if err := emitFile(filepath.Dir(output.path), filepath.Base(output.path), data); err != nil {
				return err
			}
		}
	}
	fmt.Printf("CDK service spec %s: %d mappings, %d changes, %d skipped. Report: reports/ref-db.json\n",
		pin.Version, report.ResourceCount, len(changes), len(report.Skipped))
	return nil
}
