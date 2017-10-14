package stringargv_test

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
	// Output:
	// output.json
	// input1.json input2.json
	// rule.json
	// true
}
