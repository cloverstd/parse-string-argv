[![Build Status](https://travis-ci.org/cloverstd/parse-string-argv.svg?branch=master)](https://travis-ci.org/cloverstd/parse-string-argv) [![Go Report Card](https://goreportcard.com/badge/github.com/cloverstd/parse-string-argv)](https://goreportcard.com/report/github.com/cloverstd/parse-string-argv) [![codecov](https://codecov.io/gh/cloverstd/parse-string-argv/branch/master/graph/badge.svg)](https://codecov.io/gh/cloverstd/parse-string-argv)

## Introduction

The parse-string-argv package can parse string cmd to argv list like os.Argv, and you can use [flag](https://golang.org/pkg/flag/) or other flag parse library like [pflag](https://github.com/spf13/pflag).

## Installation and usage

```bash
go get github.com/cloverstd/parse-string-argv
```

## Example

```golang
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cloverstd/parse-string-argv"
)

func Example() {
	cmd := `json -o output.json -i "input1.json input2.json" -rule-list rule.json -g`

	argv, err := stringargv.Parse(cmd)
	if err != nil {
		log.Fatalf("cannot parse cmd: %s", cmd)
	}

	output := flag.String("o", "default-output.json", "")
	input := flag.String("i", "default-input.json", "")
	ruleList := flag.String("rule-list", "default-rule-list.json", "")
	global := flag.Bool("g", false, "")

	// Be sure to remember to parse argv if you need
	// use flag.CommandLine.Parse
	// flag.Parse will parse the os.Argv
	flag.CommandLine.Parse(argv[1:])

	fmt.Println(*output)
	fmt.Println(*input)
	fmt.Println(*ruleList)
	fmt.Println(*global)
}
```

This example will generate the following output:
```bash
output.json
input1.json input2.json
rule.json
true
```
