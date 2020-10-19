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
	"os/exec"
	"path/filepath"
)

var (
	lambdaFunction string
	lambdaRole     string
	lambdaZip      string
)

func create() error {
	flag.StringVar(&lambdaFunction, "name", "", "")
	flag.StringVar(&lambdaRole, "role", "", "")
	flag.StringVar(&lambdaZip, "zip", "", "")
	flag.Parse()

	if lambdaFunction == "" {
		return errors.New("missing function name")
	}
	if lambdaRole == "" {
		return errors.New("missing role")
	}
	if lambdaZip == "" {
		lambdaZip = `fileb://` + filepath.Join(".", mainZip)
	}

	// TODO(jbd): Add other aws lambda create-function flags.
	// TODO(jbd): Check if main.zip exists.
	cmd := exec.Command("aws",
		"lambda", "create-function",
		"--function-name", lambdaFunction,
		"--runtime", "go1.x",
		"--zip-file", lambdaZip,
		"--handler", "main",
		"--role", lambdaRole,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("can't create function: %s", out)
	}
	return nil
}
