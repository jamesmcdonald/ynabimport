package skandia

import (
	"testing"

	"github.com/jamesmcdonald/ynabimport/convert"
)

func TestDateParse(t *testing.T) {
	tests := []string{
		"9999-99-99",
		"2017-07-15",
	}

	for _, test := range tests {
		if ymd := match.FindStringSubmatch(test); len(ymd) == 0 {
			t.Errorf("Failed to match valid string %s\n", test)
		}
	}
}

type TestCase struct {
	Input  string
	Output convert.Transaction
}

func TestParseLine(t *testing.T) {
	cases := []TestCase{
		TestCase{
			Input: `"2017-07-10";"2017-07-10";"12345678901";"Varekjøp";"10.07 411021 Vendor Inc Somestreet";47,00;`,
			Output: convert.Transaction{
				Date:  "2017-07-10",
				Desc:  "10.07 411021 Vendor Inc Somestreet",
				Value: "47.00",
			},
		},
		TestCase{
			Input: `"2017-07-10";"2017-07-10";"12345678901";"Varekjøp";"08.07 Kiwi 369 Øvre Åsæs 💩";665,27;`,
			Output: convert.Transaction{
				Date:  "2017-07-10",
				Desc:  "08.07 Kiwi 369 Øvre Åsæs 💩",
				Value: "665.27",
			},
		},
	}
	for _, tc := range cases {
		if result := parseLine(tc.Input); result != tc.Output {
			t.Errorf("Parse of \"%s\" produced %+v, expected %+v\n", tc.Input, result, tc.Output)
		}
	}
}

func TestParseLineWithCommaMakesDot(t *testing.T) {
	cases := []TestCase{
		TestCase{
			Input: `"2017-07-10";"2017-07-10";"12345678901";"Varekjøp";"10.07 411021 Vendor, Inc Somestreet";47,00;`,
			Output: convert.Transaction{
				Date:  "2017-07-10",
				Desc:  "10.07 411021 Vendor. Inc Somestreet",
				Value: "47.00",
			},
		},
	}
	for _, tc := range cases {
		if result := parseLine(tc.Input); result != tc.Output {
			t.Errorf("Parse of \"%s\" produced %+v, expected %+v\n", tc.Input, result, tc.Output)
		}
	}
}

func TestFormatsLoad(t *testing.T) {
	if lf := convert.ListFormats(); lf != "Skandiabanken [skandia skandiabanken]\n" &&
		lf != "Skandiabanken [skandiabanken skandia]\n" {
		t.Errorf("Incorrect formats \"%s\"\n", lf)
	}
}
