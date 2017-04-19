package ynabimport

import (
	"fmt"
	"regexp"
	"strings"
)

func init() {
	formats["DnB"] = Reader{
		encoding:  "latin1",
		parseline: dnbParseLine,
	}
	aliases["dnb"] = "DnB"
}

var dnbmatch = regexp.MustCompile(`^([0-9]{2})\.([0-9]{2})\.([0-9]{4})$`)

func dnbParseLine(source string) Transaction {
	source = strings.Replace(source, `"`, "", -1)
	source = strings.Replace(source, ",", ".", -1)
	parts := strings.Split(source, ";")
	if ymd := dnbmatch.FindStringSubmatch(parts[0]); len(ymd) > 0 {
		var value string
		if len(parts) == 5 && parts[4] != "" {
			value = "-" + parts[4]
		} else {
			value = parts[3]
		}
		t := Transaction{fmt.Sprintf("%s-%s-%s", ymd[3], ymd[2], ymd[1]), parts[1], value}
		return t
	}
	return Transaction{}
}
