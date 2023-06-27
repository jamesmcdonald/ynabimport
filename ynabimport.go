package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jamesmcdonald/ynabimport/pkg/convert"
	_ "github.com/jamesmcdonald/ynabimport/pkg/convert/bulder"
	_ "github.com/jamesmcdonald/ynabimport/pkg/convert/dnb"
	_ "github.com/jamesmcdonald/ynabimport/pkg/convert/skandia"
)

var format string
var stdout bool

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [inputfiles ...]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.StringVar(&format, "format", "skandiabanken", "the file format to import ('list' to list available formats)")
	flag.BoolVar(&stdout, "stdout", false, "write to stdout instead of matching files")
}

func main() {
	var in *os.File
	var out *os.File

	flag.Parse()

	if format == "list" {
		fmt.Print(convert.ListFormats())
		return
	}

	reader := convert.NewReader(format)
	//	if err != nil {
	//		fmt.Fprintf(stderr, "%s: could not create reader: %s\n", os.Args[0], err)
	//		os.Exit(1)
	//	}

	files := flag.Args()
	if len(files) == 0 {
		files = []string{"-"}
	}
	for _, filename := range files {
		if filename == "-" {
			in = os.Stdin
			out = os.Stdout
		} else {
			var err error
			in, err = os.Open(filename)
			if err != nil {
				panic(err)
			}
			defer in.Close()
			if stdout {
				out = os.Stdout
			} else {
				outfilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".ynabimport.csv"
				out, err = os.Create(outfilename)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s: failed to create output file %s\n",
						os.Args[0], err)
				}
				defer out.Close()
			}
		}
		reader.Process(in, out)
	}
}
