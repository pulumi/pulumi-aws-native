// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

type Reports struct {
	// UnexpectedTagsShapes is a map of the resource token of resources which have a `Tags` field but don't match an expected pattern.
	UnexpectedTagsShapes map[string]interface{}
}

func (r *Reports) WriteToDirectory(dir string) error {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("creating directory: %w", err)
	}

	writeFile := func(filename string, data any) error {
		f, err := os.Create(filepath.Join(dir, filename))
		if err != nil {
			return fmt.Errorf("creating file: %w", err)
		}
		defer contract.IgnoreClose(f)

		enc := json.NewEncoder(f)
		enc.SetIndent("", "  ")
		if err = enc.Encode(data); err != nil {
			return fmt.Errorf("writing file: %w", err)
		}
		return nil
	}

	if err = writeFile("unexpectedTagsShapes.json", r.UnexpectedTagsShapes); err != nil {
		return err
	}
	return nil
}
