// Copyright 2020 Jaana Dogan. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const mainZip = "main.zip"

func build(args ...string) error {
	dir, err := ioutil.TempDir("", "golambda")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	out := filepath.Join(dir, "main.out")
	buildArgs := []string{"build", "-o", out}
	buildArgs = append(buildArgs, args...)

	cmd := exec.Command("go", buildArgs...)
	cmd.Env = mergeEnv()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("go build failed: %s", out)
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	zipout, err := zipBinary(dir, out)
	if err != nil {
		return err
	}
	return os.Rename(zipout, filepath.Join(".", mainZip))
}

func mergeEnv() []string {
	env := os.Environ()
	return append(env, buildEnv...)
}

func zipBinary(dir, out string) (zipout string, err error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	file, err := os.Open(out)
	if err != nil {
		return "", err
	}

	f, err := w.Create("main")
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(f, file); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}

	zipout = filepath.Join(dir, mainZip)
	if err := ioutil.WriteFile(zipout, buf.Bytes(), 0644); err != nil {
		return "", err
	}
	return zipout, nil
}
