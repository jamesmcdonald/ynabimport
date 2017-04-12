package ynabimport

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestNullInput(t *testing.T) {
	var importer Importer
	importer = SkandiabankenImporter{}
	f, _ := os.Open("/dev/null")
	r, w := io.Pipe()
	defer f.Close()
	defer r.Close()
	output := make(chan []byte)
	go func() {
		data, _ := ioutil.ReadAll(r)
		output <- data
	}()
	importer.processfile(f, w)
	w.Close()
	result := string(<-output)
	if result != "Date,Payee,Category,Memo,Outflow,Inflow\n" {
		t.Errorf("Incorrect header %s\n", result)
	}
}
