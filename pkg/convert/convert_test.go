package convert_test

import (
	"bytes"
	"strings"
	"testing"

	. "github.com/jamesmcdonald/ynabimport/pkg/convert"
	_ "github.com/jamesmcdonald/ynabimport/pkg/convert/dnb"
	_ "github.com/jamesmcdonald/ynabimport/pkg/convert/skandia"
)

func TestEmptyInput(t *testing.T) {
	importer := NewReader("skandia")
	r := strings.NewReader("")
	w := new(bytes.Buffer)
	importer.Process(r, w)
	result := w.String()
	if result != "Date,Payee,Category,Memo,Outflow,Inflow\n" {
		t.Errorf("Incorrect header %s\n", result)
	}
}

func TestSkandiaInput(t *testing.T) {
	importer := NewReader("skandia")
	r := strings.NewReader("Test data\n2017-01-12;;;;;Description;;12")
	w := new(bytes.Buffer)
	importer.Process(r, w)
	result := w.String()
	if result != "Date,Payee,Category,Memo,Outflow,Inflow\n2017-01-12,Description,,,-12\n" {
		t.Errorf("Incorrect skandia result %s\n", result)
	}
}

func TestNonAliasedLookup(t *testing.T) {
	NewReader("Skandiabanken")
}

func TestFormatLoaded(t *testing.T) {
	RegisterFormat("Unaliasedformat", "utf8", nil)
	lf := strings.Split(strings.TrimSpace(ListFormats()), "\n")
	for _, format := range lf {
		if format != "Skandiabanken [skandia skandiabanken]" &&
			format != "Skandiabanken [skandiabanken skandia]" &&
			format != "DnB [dnb]" &&
			format != "Unaliasedformat" {
			t.Errorf("Incorrect format line \"%s\"\n", format)
		}
	}
}
