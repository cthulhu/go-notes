package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/cthulhu/go-notes/parser"
	"github.com/cthulhu/go-notes/scanner"
)

var usage = `Usage: go-notes [flags] <Go file or directory> ...

Without options generates all the note types. Default are:

  // FIXME    - call to fix something
  // OPTIMIZE - call for a refactoring
  // TODO     - future plans

Options:
  -f - FIXME annotations
  -o - OPTIMIZE annotations
  -t - TODO annotations
  -c CUSTOM - custom annotation label
  -format count - output format aggregated counts
  -format list - output format list with files and annotations (default)
`

var (
	fixme    = flag.Bool("f", false, "FIXME annotations")
	optimize = flag.Bool("o", false, "OPTIMIZE annotations")
	todo     = flag.Bool("t", false, "TODO annotations")
	custom   = flag.String("c", "", "custom annotation, for example BUG or BAD SMELL")
	format   = flag.String("format", "list", "output format")
)

func main() {
	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	paths, scannerErrors := scanner.New(ctx, args)
	p := parser.New(*fixme, *todo, *optimize, *custom, *format)
filesLoop:
	for {
		select {
		case file := <-paths:
			if file == "" {
				break filesLoop
			}
			exitIfError(p.Parse(file))
		case err := <-scannerErrors:
			exitIfError(err)
		}
	}
	fmt.Println(p.Aggregate())
}

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running go-notes: %v", err)
		os.Exit(1)
	}
}
