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
	"os"
	"os/exec"
	"path/filepath"
)

func create() error {
	// TODO(jbd): Add other aws lambda create-function flags.
	var name, role, zip, endpointURL string
	fset := flag.NewFlagSet("create", flag.ExitOnError)
	fset.StringVar(&name, "name", "", "")
	fset.StringVar(&role, "role", "", "")
	fset.StringVar(&zip, "zip", "", "")
	fset.StringVar(&endpointURL, "endpoint-url", "", "")
	fset.Parse(os.Args[2:])

	if name == "" {
		return errors.New("missing function name")
	}
	if role == "" {
		return errors.New("missing role")
	}
	if zip == "" {
		zip = defaultZip()
	}

	args := []string{"lambda"}
	if endpointURL != "" {
		args = append(args, "--endpoint-url", endpointURL)
	}
	// TODO(jbd): Check if main.zip exists.
	args = append(args, "create-function",
		"--function-name", name,
		"--runtime", "go1.x",
		"--zip-file", zip,
		"--handler", "main",
		"--role", role,
	)
	cmd := exec.Command("aws", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func update() error {
	var name, zip, endpointURL string
	var publish bool
	fset := flag.NewFlagSet("update", flag.ExitOnError)
	fset.StringVar(&name, "name", "", "")
	fset.BoolVar(&publish, "publish", true, "")
	fset.StringVar(&zip, "zip", "", "")
	fset.StringVar(&endpointURL, "endpoint-url", "", "")
	fset.Parse(os.Args[2:])

	if name == "" {
		return errors.New("missing function name")
	}
	if zip == "" {
		zip = defaultZip()
	}
	args := []string{"lambda"}
	if endpointURL != "" {
		args = append(args, "--endpoint-url", endpointURL)
	}
	args = append(args, "update-function-code",
		"--function-name", name,
		"--zip-file", zip,
	)
	cmd := exec.Command("aws", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func defaultZip() string {
	return `fileb://` + filepath.Join(".", mainZip)
}
