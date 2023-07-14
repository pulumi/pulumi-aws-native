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
