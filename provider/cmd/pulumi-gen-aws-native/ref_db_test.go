// Copyright 2026, Pulumi Corporation.

//nolint:goconst // Repeated literals keep test and schema fixtures readable.
package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi-aws-native/provider/pkg/metadata"
	"github.com/pulumi/pulumi-aws-native/provider/pkg/refdb"
)

const refDBFixture = `{
  "schema": {"resource": {"type": "entities", "entities": [
    {"cloudFormationType": "Explicit", "primaryIdentifier": ["Parent", "Name"],
     "cfnRefIdentifier": ["Arn"], "properties": {"Parent": {}, "Name": {}}, "attributes": {"Arn": {}}},
    {"cloudFormationType": "Composite", "primaryIdentifier": ["Parent", "Name"],
     "properties": {"Parent": {}, "Name": {}}, "attributes": {}},
    {"cloudFormationType": "Nested", "primaryIdentifier": ["Config/Id"],
     "properties": {"Config": {}}, "attributes": {}},
    {"cloudFormationType": "Missing", "properties": {}, "attributes": {}},
    {"cloudFormationType": "Unknown", "primaryIdentifier": ["Id"], "properties": {}, "attributes": {}},
    {"cloudFormationType": "EmptyRef", "primaryIdentifier": ["Name"],
     "cfnRefIdentifier": [], "properties": {"Name": {}}, "attributes": {}}
  ]}}
}`

func testServiceSpecDatabase(t *testing.T, data string) serviceSpecDatabase {
	t.Helper()
	var database serviceSpecDatabase
	require.NoError(t, json.Unmarshal([]byte(data), &database))
	return database
}

func testServiceSpecArchive(t *testing.T, data string, filename string) ([]byte, string) {
	t.Helper()
	var database bytes.Buffer
	compressedDB := gzip.NewWriter(&database)
	_, err := io.WriteString(compressedDB, data)
	require.NoError(t, err)
	require.NoError(t, compressedDB.Close())
	var archive bytes.Buffer
	compressedArchive := gzip.NewWriter(&archive)
	entries := tar.NewWriter(compressedArchive)
	require.NoError(t, entries.WriteHeader(&tar.Header{
		Name: filename, Mode: 0o600, Size: int64(database.Len()), Typeflag: tar.TypeReg,
	}))
	_, err = entries.Write(database.Bytes())
	require.NoError(t, err)
	require.NoError(t, entries.Close())
	require.NoError(t, compressedArchive.Close())
	digest := sha512.Sum512(archive.Bytes())
	return archive.Bytes(), "sha512-" + base64.StdEncoding.EncodeToString(digest[:])
}

func TestRefDBTranslation(t *testing.T) {
	result, report, err := generateRefDB(testServiceSpecDatabase(t, refDBFixture), nil)
	require.NoError(t, err)
	require.Len(t, result.Resources, 2)
	assert.Equal(t, metadata.CfRefBehavior{Property: "Arn"}, result.Resources["Explicit"].RefReturns.CfRefBehavior)
	assert.Equal(t, metadata.CfRefBehavior{Properties: []string{"Parent", "Name"}, Delimiter: "|"},
		result.Resources["Composite"].RefReturns.CfRefBehavior)
	assert.Equal(t, map[string]string{
		"Nested":   "Identifiers unavailable as top-level properties: Config/Id",
		"Unknown":  "Identifiers unavailable as top-level properties: Id",
		"Missing":  "No Ref or primary identifier",
		"EmptyRef": "No Ref or primary identifier",
	}, report.Skipped)
	assert.Empty(t, report.Overrides)
}

func TestRefDBOverrides(t *testing.T) {
	overrides := map[string]refDBOverride{
		"Explicit": {Reason: "Verified CF exception", RefReturns: metadata.CfRefBehavior{Property: "Name"}},
		"Missing":  {Reason: "Verified unsupported Ref", RefReturns: metadata.CfRefBehavior{NotSupported: true}},
		"Absent":   {Reason: "Verified legacy resource", RefReturns: metadata.CfRefBehavior{Property: "Id"}},
	}
	result, report, err := generateRefDB(testServiceSpecDatabase(t, refDBFixture), overrides)
	require.NoError(t, err)
	assert.Equal(t, "Name", result.Resources["Explicit"].RefReturns.Property)
	assert.True(t, result.Resources["Missing"].RefReturns.NotSupported)
	assert.Equal(t, "Id", result.Resources["Absent"].RefReturns.Property)
	assert.NotContains(t, report.Skipped, "Missing")
	require.NotNil(t, report.Overrides["Explicit"].Upstream)
	assert.Equal(t, "Arn", report.Overrides["Explicit"].Upstream.Property)
	assert.Nil(t, report.Overrides["Absent"].Upstream)
	for name, override := range map[string]refDBOverride{
		"no reason":   {RefReturns: metadata.CfRefBehavior{Property: "Name"}},
		"no behavior": {Reason: "Evidence"},
		"conflicting behaviors": {
			Reason: "Evidence", RefReturns: metadata.CfRefBehavior{Property: "Name", NotSupported: true},
		},
		"no delimiter": {Reason: "Evidence", RefReturns: metadata.CfRefBehavior{Properties: []string{"A", "B"}}},
	} {
		t.Run(name, func(t *testing.T) {
			_, _, err := generateRefDB(testServiceSpecDatabase(t, refDBFixture), map[string]refDBOverride{"Explicit": override})
			require.Error(t, err)
		})
	}
}

func TestRefDBRejectsBrokenUpstream(t *testing.T) {
	for name, data := range map[string]string{
		"changed format": `{"resources": []}`,
		"wrong table":    `{"schema":{"resource":{"type":"other","entities":[{}]}}}`,
		"empty table":    `{"schema":{"resource":{"type":"entities","entities":[]}}}`,
		"missing name":   `{"schema":{"resource":{"type":"entities","entities":[{}]}}}`,
		"duplicate name": `{"schema":{"resource":{"type":"entities","entities":[
			{"cloudFormationType":"A"},{"cloudFormationType":"A"}]}}}`,
		"empty identifier": `{"schema":{"resource":{"type":"entities","entities":[
			{"cloudFormationType":"A","primaryIdentifier":[""],"properties":{},"attributes":{}}]}}}`,
	} {
		t.Run(name, func(t *testing.T) {
			_, _, err := generateRefDB(testServiceSpecDatabase(t, data), map[string]refDBOverride{
				"Legacy": {Reason: "Evidence", RefReturns: metadata.CfRefBehavior{Property: "Name"}},
			})
			require.Error(t, err)
		})
	}
}

func TestRefDBRejectsMissingPropertyMaps(t *testing.T) {
	database := testServiceSpecDatabase(t, refDBFixture)
	database.Schema.Resource.Entities[0].Properties = nil
	_, _, err := generateRefDB(database, nil)
	require.ErrorContains(t, err, "missing properties or attributes in CDK resource Explicit")
}

func TestServiceSpecArchive(t *testing.T) {
	archive, integrity := testServiceSpecArchive(t, refDBFixture, "package/db.json.gz")
	database, err := decodeServiceSpecArchive(archive, integrity)
	require.NoError(t, err)
	assert.Equal(t, testServiceSpecDatabase(t, refDBFixture), database)
	_, err = decodeServiceSpecArchive([]byte("corrupt archive"), integrity)
	require.ErrorContains(t, err, "integrity mismatch")
	_, err = decodeServiceSpecArchive(archive, "sha512-wrong")
	require.ErrorContains(t, err, "integrity mismatch")
	archive, integrity = testServiceSpecArchive(t, refDBFixture, "other.json.gz")
	_, err = decodeServiceSpecArchive(archive, integrity)
	require.ErrorContains(t, err, "no package/db.json.gz")
}

func TestServiceSpecHTTPFailures(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "unavailable", http.StatusServiceUnavailable)
	}))
	defer server.Close()
	source := serviceSpecSource{client: server.Client(), registry: server.URL}
	_, err := source.latest(t.Context())
	require.ErrorContains(t, err, "503 Service Unavailable")
	_, err = source.load(t.Context(), serviceSpecPin{Version: "../../invalid"})
	require.ErrorContains(t, err, "invalid service spec version")
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	_, err = source.latest(ctx)
	require.ErrorIs(t, err, context.Canceled)
}

func testRefDBWorkspace(t *testing.T, integrity string) string {
	t.Helper()
	root := t.TempDir()
	require.NoError(t, os.Mkdir(filepath.Join(root, "meta"), 0o700))
	require.NoError(t, emitFile(filepath.Join(root, "meta"), "service-spec-version.json",
		[]byte(`{"version":"0.1.1","integrity":"`+integrity+`"}`)))
	require.NoError(t, emitFile(filepath.Join(root, "meta"), "ref-db-overrides.json", []byte(`{"resources":{}}`)))
	require.NoError(t, emitFile(filepath.Join(root, "meta"), "ref-db.json", []byte(`{"resources":{}}`)))
	return root
}

func TestRefDBGenerationAndUpdate(t *testing.T) {
	archive, integrity := testServiceSpecArchive(t, refDBFixture, "package/db.json.gz")
	var requests []string
	var requestsMu sync.Mutex
	takeRequests := func() []string {
		requestsMu.Lock()
		defer requestsMu.Unlock()
		result := requests
		requests = nil
		return result
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestsMu.Lock()
		requests = append(requests, r.URL.Path)
		requestsMu.Unlock()
		switch r.URL.Path {
		case "/latest":
			_, err := io.WriteString(w, `{"version":"0.1.2","dist":{"integrity":"`+integrity+`"}}`)
			assert.NoError(t, err)
		case "/-/aws-service-spec-0.1.1.tgz", "/-/aws-service-spec-0.1.2.tgz":
			_, err := w.Write(archive)
			assert.NoError(t, err)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()
	source := serviceSpecSource{client: server.Client(), registry: server.URL}
	root := testRefDBWorkspace(t, integrity)
	require.NoError(t, runRefDB(t.Context(), root, refDBOptions{}, source))
	assert.Equal(t, []string{"/-/aws-service-spec-0.1.1.tgz"}, takeRequests())
	target := filepath.Join(root, "meta", "ref-db.json")
	reportPath := filepath.Join(root, "reports", "ref-db.json")
	expected, err := os.ReadFile(target)
	require.NoError(t, err)
	expectedReport, err := os.ReadFile(reportPath)
	require.NoError(t, err)
	require.NoError(t, runRefDB(t.Context(), root, refDBOptions{}, source))
	actual, err := os.ReadFile(target)
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
	actualReport, err := os.ReadFile(reportPath)
	require.NoError(t, err)
	assert.Equal(t, expectedReport, actualReport)
	require.NoError(t, runRefDB(t.Context(), root, refDBOptions{Check: true}, source))

	// A stale check must not rewrite either artifact.
	require.NoError(t, os.WriteFile(target, []byte(`{"resources":{}}`), 0o600))
	require.ErrorContains(t, runRefDB(t.Context(), root, refDBOptions{Check: true}, source), "stale")
	actual, err = os.ReadFile(target)
	require.NoError(t, err)
	assert.Equal(t, `{"resources":{}}`, string(actual))
	actualReport, err = os.ReadFile(reportPath)
	require.NoError(t, err)
	assert.Equal(t, expectedReport, actualReport)
	require.NoError(t, os.WriteFile(target, expected, 0o600)) //nolint:gosec // target is confined to t.TempDir.
	require.NoError(t, os.WriteFile(reportPath, []byte(`{}`), 0o600))
	require.ErrorContains(t, runRefDB(t.Context(), root, refDBOptions{Check: true}, source), "stale")

	_ = takeRequests()
	require.NoError(t, runRefDB(t.Context(), root, refDBOptions{Update: true}, source))
	assert.Equal(t, []string{"/latest", "/-/aws-service-spec-0.1.2.tgz"}, takeRequests())
	var pin serviceSpecPin
	require.NoError(t, readRefDBJSON(filepath.Join(root, "meta", "service-spec-version.json"), &pin))
	assert.Equal(t, serviceSpecPin{Version: "0.1.2", Integrity: integrity}, pin)
	require.NoError(t, runRefDB(t.Context(), root, refDBOptions{Check: true}, source))
}

func TestRefDBFailedUpdateDoesNotWrite(t *testing.T) {
	root := testRefDBWorkspace(t, "old-integrity")
	target := filepath.Join(root, "meta", "ref-db.json")
	previous := []byte(`{"resources":{"Existing":{"refReturns":{"property":"Name"}}}}`)
	require.NoError(t, os.WriteFile(target, previous, 0o600))
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/latest" {
			_, err := io.WriteString(w, `{"version":"0.1.2","dist":{"integrity":"sha512-invalid"}}`)
			assert.NoError(t, err)
		} else {
			_, err := io.WriteString(w, "corrupt")
			assert.NoError(t, err)
		}
	}))
	defer server.Close()
	source := serviceSpecSource{client: server.Client(), registry: server.URL}
	err := runRefDB(t.Context(), root, refDBOptions{Update: true}, source)
	require.ErrorContains(t, err, "integrity mismatch")
	var pin serviceSpecPin
	require.NoError(t, readRefDBJSON(filepath.Join(root, "meta", "service-spec-version.json"), &pin))
	assert.Equal(t, serviceSpecPin{Version: "0.1.1", Integrity: "old-integrity"}, pin)
	actual, err := os.ReadFile(target)
	require.NoError(t, err)
	assert.Equal(t, previous, actual)
	assert.NoFileExists(t, filepath.Join(root, "reports", "ref-db.json"))
}

func TestRefDBOfflineAndInvalidOptions(t *testing.T) {
	root := testRefDBWorkspace(t, "unused")
	fixture := filepath.Join(root, "db.json")
	require.NoError(t, os.WriteFile(fixture, []byte(refDBFixture), 0o600))
	// No HTTP client is provided. Offline generation must not use the network.
	require.NoError(t, runRefDB(t.Context(), root, refDBOptions{Database: fixture}, serviceSpecSource{}))
	for _, options := range []refDBOptions{{Update: true, Check: true}, {Update: true, Database: fixture}} {
		require.ErrorContains(t, runRefDB(t.Context(), root, options, serviceSpecSource{}), "cannot be combined")
	}
	require.NoError(t, os.WriteFile(filepath.Join(root, "meta", "ref-db-overrides.json"),
		[]byte(`{"resources":{"Explicit":{"reason":"Evidence","refReturns":{"typo":"Name"}}}}`), 0o600))
	err := runRefDB(t.Context(), root, refDBOptions{Database: fixture}, serviceSpecSource{})
	require.ErrorContains(t, err, "unknown field")
}

func TestRefDBChangeReport(t *testing.T) {
	var previous, current refdb.RefDB
	require.NoError(t, json.Unmarshal([]byte(`{"resources":{
		"Same":{"refReturns":{"property":"Name","heuristic":"guess"}},
		"Changed":{"refReturns":{"property":"Name"}},
		"Removed":{"refReturns":{"property":"Id"}}
	}}`), &previous))
	require.NoError(t, json.Unmarshal([]byte(`{"resources":{
		"Same":{"refReturns":{"property":"Name"}},
		"Changed":{"refReturns":{"properties":["Parent","Name"],"delimiter":"|"}},
		"Added":{"refReturns":{"property":"Arn"}}
	}}`), &current))
	changes := compareRefDB(&previous, &current)
	assert.NotContains(t, changes, "Same")
	require.Len(t, changes, 3)
	assert.Equal(t, &metadata.CfRefBehavior{Property: "Name"}, changes["Changed"].Before)
	assert.Equal(t, &metadata.CfRefBehavior{Properties: []string{"Parent", "Name"}, Delimiter: "|"},
		changes["Changed"].After)
	assert.Nil(t, changes["Removed"].After)
	assert.Nil(t, changes["Added"].Before)
}
