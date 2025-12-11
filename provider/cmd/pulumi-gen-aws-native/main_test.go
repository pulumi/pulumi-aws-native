// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"io"
	"os"
	"testing"

	jsschema "github.com/pulumi/jsschema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCleanDir(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	//create tempdir
	temp, err := os.MkdirTemp("", "test_CleanDir")
	require.Nil(err)
	t.Cleanup(func() { os.RemoveAll(temp) })

	dir := temp + "/foo"

	// when called on non-existent dir
	require.NoDirExists(dir)
	err = cleanDir(dir, 0755)
	if assert.Nil(err) {
		assertExistsAndEmpty(t, dir)
	}

	// when called on an existing Dir
	require.DirExists(dir)
	file, err := os.Create(dir + "/temp.txt")
	require.Nil(err)
	file.Close()

	err = cleanDir(dir, 0755)
	if assert.Nil(err) {
		assertExistsAndEmpty(t, dir)
	}
}

func Test_mergeAutoNaming(t *testing.T) {
	t.Run("overlay overwrites", func(t *testing.T) {
		schema := &jsschema.Schema{
			Extras: map[string]interface{}{
				"typeName": "AWS::IAM::Role",
			},
			Properties: map[string]*jsschema.Schema{
				"RoleName": {},
			},
		}

		overlay := map[string]AutoNamingOverlay{
			"AWS::IAM::Role": {
				"RoleName": {
					MaxLength: 64,
					MinLength: 1,
				},
			},
		}

		mergeAutoNaming(schema, overlay)

		assert.Equal(t, &jsschema.Schema{
			MaxLength: jsschema.Integer{Val: 64},
			MinLength: jsschema.Integer{Val: 1},
		}, schema.Properties["RoleName"])
	})

	t.Run("overlay does not overwrite", func(t *testing.T) {
		schema := &jsschema.Schema{
			Extras: map[string]interface{}{
				"typeName": "AWS::IAM::Role",
			},
			Properties: map[string]*jsschema.Schema{
				"RoleName": {
					MaxLength: jsschema.Integer{Val: 100},
					MinLength: jsschema.Integer{Val: 4},
				},
			},
		}

		overlay := map[string]AutoNamingOverlay{
			"AWS::IAM::Role": {
				"RoleName": {
					MaxLength: 64,
					MinLength: 1,
				},
			},
		}

		mergeAutoNaming(schema, overlay)

		assert.Equal(t, &jsschema.Schema{
			MaxLength: jsschema.Integer{Val: 100},
			MinLength: jsschema.Integer{Val: 4},
		}, schema.Properties["RoleName"])
	})

	t.Run("overlay not found", func(t *testing.T) {
		schema := &jsschema.Schema{
			Extras: map[string]interface{}{
				"typeName": "AWS::IAM::Role",
			},
			Properties: map[string]*jsschema.Schema{
				"RoleName": {},
			},
		}

		overlay := map[string]AutoNamingOverlay{}

		mergeAutoNaming(schema, overlay)

		assert.Equal(t, &jsschema.Schema{
			MaxLength: jsschema.Integer{Val: 0},
			MinLength: jsschema.Integer{Val: 0},
		}, schema.Properties["RoleName"])
	})
}

func Test_readJsonSchema_convertsExclusiveBounds(t *testing.T) {
	tempDir := t.TempDir()
	schemaPath := tempDir + "/schema.json"
	err := os.WriteFile(schemaPath, []byte(`
{
  "typeName": "AWS::Test::Thing",
  "properties": {
    "Size": {
      "type": "integer",
      "exclusiveMaximum": 10
    }
  }
}`), 0o600)
	require.NoError(t, err)

	schema := readJsonSchema(schemaPath, nil)
	size := schema.Properties["Size"]
	require.NotNil(t, size)
	assert.Equal(t, float64(10), size.Maximum.Val)
	assert.True(t, size.ExclusiveMaximum.Bool())
}

func assertExistsAndEmpty(t *testing.T, path string) bool {
	assert := assert.New(t)

	if !assert.DirExists(path) {
		return false
	}

	dir, err := os.Open(path)
	if !assert.Nil(err) {
		return false
	}
	defer dir.Close()

	_, err = dir.Readdirnames(1)
	return assert.Equalf(io.EOF, err, "Non-empty dir")
}
