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
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		printUsage(1)
	}
	flag.Usage = func() {
		printUsage(1)
	}
	flag.Parse()

	switch os.Args[1] {
	case "init":
		// TODO
	case "build":
		var pkg string
		if len(os.Args) == 2 {
			pkg = "."
		} else {
			pkg = os.Args[2]
		}
		if err := build(pkg); err != nil {
			log.Fatal(err)
		}
	}
}

func build(pkg string) error {
	dir, err := ioutil.TempDir("", "golambda")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	out := filepath.Join(dir, "main.out")
	cmd := exec.Command("go", "build", "-o", out, "-v")
	cmd.Env = mergeEnv()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("go build failed: %s", out)
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	// TODO(jbd): Pass go build flags to go build.
	zipout, err := zipBinary(dir, out)
	if err != nil {
		return err
	}
	return os.Rename(zipout, filepath.Join(".", "main.zip"))
}

func mergeEnv() []string {
	env := os.Environ()
	return append(env, buildEnv...)
}

func printUsage(code int) {
	fmt.Print(usageText)
	os.Exit(code)
}

const usageText = `golambda [cmd] <options>

Commands:
  - build  Builds the package and generates a zip.
`

// TODO(jbd): Add init and deploy subcommands.
