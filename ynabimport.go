package main

import (
	"bufio"
	"fmt"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type transaction struct {
	date  string
	desc  string
	value string
}

func (t *transaction) String() string {
	return fmt.Sprintf("%s,%s,,,%s", t.date, t.desc, t.value)
}

var match = regexp.MustCompile(`^([0-9]{4})-([0-9]{2})-([0-9]{2})$`)

func parseline(source string) transaction {
	source = strings.Replace(source, `"`, "", -1)
	source = strings.Replace(source, ",", ".", -1)
	parts := strings.Split(source, ";")
	if ymd := match.FindStringSubmatch(parts[0]); len(ymd) > 0 {
		var value string
		if len(parts) == 7 && parts[6] != "" {
			value = "-" + parts[6]
		} else {
			value = parts[5]
		}
		t := transaction{fmt.Sprintf("%s-%s-%s", ymd[1], ymd[2], ymd[3]), parts[4], value}
		return t
	}
	return transaction{}
}

func processfile(rawin io.Reader, rawout io.Writer) {
	out := bufio.NewWriter(rawout)
	r, _ := charset.NewReader("latin1", rawin)
	scanner := bufio.NewScanner(r)
	var ledger []transaction
	for scanner.Scan() {
		if t := parseline(scanner.Text()); len(t.value) > 0 {
			ledger = append(ledger, t)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Fprintf(out, "%s\n", "Date,Payee,Category,Memo,Outflow,Inflow")
	for _, t := range ledger {
		fmt.Fprintf(out, "%s\n", t.String())
	}
	out.Flush()
}

func main() {
	var in *os.File
	var out *os.File
	for _, filename := range os.Args[1:] {
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
			outfilename := strings.TrimSuffix(filename, filepath.Ext(filename)) + ".ynabimport.csv"
			out, err = os.Create(outfilename)
			if err != nil {
				panic(err)
			}
			defer out.Close()
		}
		processfile(in, out)
	}
}

