package ynabimport

import (
	"fmt"
	"regexp"
	"strings"
)

func init() {
	formats["Skandiabanken"] = Reader{
		encoding:  "latin1",
		parseline: SkandiaParseLine,
	}
	aliases["skandiabanken"] = "Skandiabanken"
	aliases["skandia"] = "Skandiabanken"
}

type SkandiaFormat struct {
	encoding string
}

var match = regexp.MustCompile(`^([0-9]{4})-([0-9]{2})-([0-9]{2})$`)

func SkandiaParseLine(source string) Transaction {
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
