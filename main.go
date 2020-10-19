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
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage(1)
	}
	flag.Usage = func() {
		printUsage(1)
	}

	switch os.Args[1] {
	case "init":
		// TODO(jbd): Implement.
	case "build":
		if err := build(os.Args[2:]...); err != nil {
			log.Fatal(err)
		}
	case "create":
		if err := create(); err != nil {
			log.Fatal(err)
		}
	case "update":
		// TODO(jbd): Implement.
	}
}

func printUsage(code int) {
	fmt.Print(usageText)
	os.Exit(code)
}

const usageText = `golambda [cmd] <options>

Commands:
  - build  Builds the package and generates a zip.
`

// TODO(jbd): Add init, create and update subcommands.
