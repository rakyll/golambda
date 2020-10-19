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
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func create() error {
	// TODO(jbd): Add other aws lambda create-function flags.
	var name, role, zip string
	fset := flag.NewFlagSet("create", flag.ExitOnError)
	fset.StringVar(&name, "name", "", "")
	fset.StringVar(&role, "role", "", "")
	fset.StringVar(&zip, "zip", "", "")
	fset.Parse(os.Args[2:])

	if name == "" {
		return errors.New("missing function name")
	}
	if role == "" {
		return errors.New("missing role")
	}
	if zip == "" {
		zip = `fileb://` + filepath.Join(".", mainZip)
	}

	// TODO(jbd): Check if main.zip exists.
	cmd := exec.Command("aws",
		"lambda", "create-function",
		"--function-name", name,
		"--runtime", "go1.x",
		"--zip-file", zip,
		"--handler", "main",
		"--role", role,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("can't create function: %s", out)
	}
	return nil
}
