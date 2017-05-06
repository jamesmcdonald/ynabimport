package dnb

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jamesmcdonald/ynabimport"
)

func init() {
	ynabimport.RegisterFormat("DnB", "latin1", parseLine)
	ynabimport.RegisterAlias("dnb", "DnB")
}

var match = regexp.MustCompile(`^([0-9]{2})\.([0-9]{2})\.([0-9]{4})$`)

func parseLine(source string) ynabimport.Transaction {
	source = strings.Replace(source, `"`, "", -1)
	source = strings.Replace(source, ",", ".", -1)
	parts := strings.Split(source, ";")
	if ymd := match.FindStringSubmatch(parts[0]); len(ymd) > 0 {
		var value string
		if len(parts) == 5 && parts[4] != "" {
			value = "-" + parts[4]
		} else {
			value = parts[3]
		}
		t := ynabimport.Transaction{fmt.Sprintf("%s-%s-%s", ymd[3], ymd[2], ymd[1]), parts[1], value}
		return t
	}
	return ynabimport.Transaction{}
}
