package ynabimport

import (
	"bytes"
	"strings"
	"testing"
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
	r := strings.NewReader("Test data\n2017-01-12;;;;Description;;12")
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
