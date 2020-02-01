package dnb

import (
	"testing"

	"github.com/jamesmcdonald/ynabimport/pkg/convert"
)

func TestMatch(t *testing.T) {
	tests := []string{
		"99.99.9999",
		"15.07.2017",
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
			Input: `"10.07.2017";"10.07 411021 Vendor Inc Somestreet";;47,00;`,
			Output: convert.Transaction{
				Date:  "2017-07-10",
				Desc:  "10.07 411021 Vendor Inc Somestreet",
				Value: "47.00",
			},
		},
		TestCase{
			Input: `"10.07.2017";"08.07 Kiwi 369 Øvre Åsæs 💩";;665,27;`,
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
			Input: `"10.07.2017";"08.07 Kiwi, 369 Øvre Åsæs 💩";;665,27;`,
			Output: convert.Transaction{
				Date:  "2017-07-10",
				Desc:  "08.07 Kiwi. 369 Øvre Åsæs 💩",
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

func TestFormatsLoad(t *testing.T) {
	if lf := convert.ListFormats(); lf != "DnB [dnb]\n" {
		t.Errorf("Incorrect formats \"%s\"\n", lf)
	}
}
