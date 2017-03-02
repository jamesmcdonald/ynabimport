package main

import (
	"bufio"
	"fmt"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
	"os"
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

func main() {
	r, _ := charset.NewReader("latin1", os.Stdin)
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
	fmt.Println("Date,Payee,Category,Memo,Outflow,Inflow")
	for _, t := range ledger {
		fmt.Println(t.String())
	}
}
