package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/twmb/chkjson"
)

func main() {
	var escHTML = flag.Bool("html", false, "ensure the escaping will be safe for HTML")
	var escJSONP = flag.Bool("jsonp", false, "ensure the escaping will be safe for JSONP")
	flag.Usage = func() {
		fmt.Print(`jsonesc - consume STDIN and print it JSON string escaped

Usage:
  jsonesc [-html] [-jsonp] STDIN

Flags:
`)
		flag.PrintDefaults()
	}
	flag.Parse()
	var opts []chkjson.EscapeOpt
	if *escHTML {
		opts = append(opts, chkjson.EscapeHTML)
	}
	if *escJSONP {
		opts = append(opts, chkjson.EscapeJSONP)
	}
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("unable to read stdin: %v", err)
		os.Exit(1)
	}
	out := make([]byte, 0, 5*len(in)/4)
	out = chkjson.Escape(out, in, opts...)
	os.Stdout.Write(out)
}
