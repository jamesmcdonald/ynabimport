package main

import (
	"flag"
	"fmt"
	"github.com/jamesmcdonald/ynabimport"
	"os"
	"path/filepath"
	"strings"
)

var format string
var stdout bool

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [inputfiles ...]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.StringVar(&format, "format", "skandiabanken", "the file format to import")
	flag.BoolVar(&stdout, "stdout", false, "write to stdout instead of matching files")
}

func main() {
	var in *os.File
	var out *os.File

	flag.Parse()

	importer, err := ynab.NewImporter(format)
	if err != nil {
		fmt.Fprintf(stderr, "%s: could not create importer: %s\n", os.Args[0], err)
		os.Exit(1)
	}

	switch format {
	case "skandiabanken":
		importer = ynabimport.SkandiabankenImporter{}
	default:
		fmt.Fprintf(os.Stderr, "unknown format %s\n", format)
		os.Exit(2)
	}

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
		importer.Processfile(in, out)
	}
}
