package ynabimport

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
)

type Transaction struct {
	Date  string
	Desc  string
	Value string
}

func (t Transaction) String() string {
	return fmt.Sprintf("%s,%s,,,%s", t.Date, t.Desc, t.Value)
}

type Reader struct {
	encoding  string
	parseline func(source string) Transaction
}

// Send each line to the reader's line parser. Print all valid transactions.
func (reader *Reader) Process(rawin io.Reader, rawout io.Writer) {
	out := bufio.NewWriter(rawout)
	r, _ := charset.NewReader("latin1", rawin)
	scanner := bufio.NewScanner(r)
	var ledger Ledger
	for scanner.Scan() {
		if t := reader.parseline(scanner.Text()); len(t.Value) > 0 {
			ledger = append(ledger, t)
		}
	}
	if err := scanner.Err(); err != nil {
		// TODO decide what should happen here
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Fprint(out, ledger)
	out.Flush()
}

type Ledger []Transaction

func (ledger Ledger) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Date,Payee,Category,Memo,Outflow,Inflow\n")
	for _, t := range ledger {
		fmt.Fprintf(&buffer, "%s\n", t)
	}
	return buffer.String()
}

var formats map[string]Reader = make(map[string]Reader)
var aliases map[string]string = make(map[string]string)

func RegisterFormat(name string, encoding string, parseline func(string) Transaction) {
	formats[name] = Reader{encoding, parseline}
}

func RegisterAlias(alias string, name string) {
	aliases[alias] = name
}

func NewReader(alias string) Reader {
	name, ok := aliases[alias]
	if !ok {
		name = alias
	}
	return formats[name]
}

func ListFormats() string {
	var buf bytes.Buffer
	var formatAliases []string
	for format := range formats {
		formatAliases = []string{}
		for alias, name := range aliases {
			if name == format && alias != format {
				formatAliases = append(formatAliases, alias)
			}
		}
		if len(formatAliases) > 0 {
			fmt.Fprintf(&buf, "%s %v\n", format, formatAliases)
		} else {
			fmt.Fprintf(&buf, "%s\n", format)
		}
	}
	return buf.String()
}
