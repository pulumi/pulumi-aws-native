// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func CompressSchema(pkgSpec schema.PackageSpec) ([]byte, error) {
	compressedSchema := bytes.Buffer{}
	compressedWriter := gzip.NewWriter(&compressedSchema)
	err := json.NewEncoder(compressedWriter).Encode(pkgSpec)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling schema")
	}
	if err = compressedWriter.Close(); err != nil {
		return nil, err
	}
	return compressedSchema.Bytes(), nil
}

func DecompressSchema(compressedSchema []byte) ([]byte, error) {
	uncompressed, err := gzip.NewReader(bytes.NewReader(compressedSchema))
	if err != nil {
		return nil, errors.Wrap(err, "expand compressed schema")
	}
	uncompressedBuf := bytes.Buffer{}
	if _, err = io.Copy(&uncompressedBuf, uncompressed); err != nil {
		return nil, err
	}
	if err = uncompressed.Close(); err != nil {
		return nil, errors.Wrap(err, "closing uncompress stream for schema")
	}
	return uncompressedBuf.Bytes(), err
}
