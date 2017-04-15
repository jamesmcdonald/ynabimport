package ynabimport

import (
	"bufio"
	"fmt"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
	"io"
	"os"
	"regexp"
	"strings"
)

type Transaction struct {
	date  string
	desc  string
	value string
}

type Ledger []Transaction

func (ledger *Ledger) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Date,Payee,Category,Memo,Outflow,Inflow\n")
	for _, t := range ledger {
		buffer.WriteString(t)
		buffer.WriteRune('\n')
	}
	return buffer.String()
}

type Reader interface {
	Read() (t Transaction)
	ReadAll() (t []Transaction)
}

var formats map[string]Format
var aliases map[string]string

func NewImporter(format string) (Importer, error) {
	i, exists := importers[format]
	if !exists {
		return nil, Errorf("unknown format %s", format)
	}
	i := Importer{}
	i.encoding = importers[format].encoding
}

func init() {
	formats = map[string]string{
		"Skandiabanken": &SkandiaFormat{},
	}
}

func (t *Transaction) String() string {
	return fmt.Sprintf("%s,%s,,,%s", t.date, t.desc, t.value)
}

type SkandiaFormat struct {
	encoding string
}

var match = regexp.MustCompile(`^([0-9]{4})-([0-9]{2})-([0-9]{2})$`)

func parseline(source string) Transaction {
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
		t := Transaction{fmt.Sprintf("%s-%s-%s", ymd[1], ymd[2], ymd[3]), parts[4], value}
		return t
	}
	return Transaction{}
}

func (format SkandiaFormat) Processfile(rawin io.Reader, rawout io.Writer) {
	out := bufio.NewWriter(rawout)
	r, _ := charset.NewReader(format.encoding, rawin)
	scanner := bufio.NewScanner(r)
	var ledger Ledger
	for scanner.Scan() {
		if t := parseline(scanner.Text()); len(t.value) > 0 {
			ledger = append(ledger, t)
		}
	}
	if err := scanner.Err(); err != nil {
		// TODO decide what should happen here
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Fprint(out, ledger)
	out.Flush()
}

func Blah(rawin io.Reader, rawout io.Writer, formatName string) {
	format := makeMeAFormat(formatName)
	format.Processfile(rawin, rawout)
}
