package skandia

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jamesmcdonald/ynabimport"
)

func init() {
	ynabimport.RegisterFormat("Skandiabanken", "latin1", parseLine)
	ynabimport.RegisterAlias("skandia", "Skandiabanken")
	ynabimport.RegisterAlias("skandiabanken", "Skandiabanken")
}

var match = regexp.MustCompile(`^([0-9]{4})-([0-9]{2})-([0-9]{2})$`)

func parseLine(source string) ynabimport.Transaction {
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
		t := ynabimport.Transaction{fmt.Sprintf("%s-%s-%s", ymd[1], ymd[2], ymd[3]), parts[4], value}
		return t
	}
	return ynabimport.Transaction{}
}
