package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldMatchDay(t *testing.T) {
	tests := []struct {
		name  string
		input string
		match bool
	}{
		{
			name:  "monday with date",
			input: "Poniedziałek 23.10",
			match: true,
		},
		{
			name:  "tuesday with date",
			input: "Wtorek 23.10",
			match: true,
		},
		{
			name:  "wednesday with date",
			input: "Środa 23.10",
			match: true,
		},
		{
			name:  "thursday with date",
			input: "Czwartek 23.10",
			match: true,
		},
		{
			name:  "friday with date",
			input: "Piątek 23.10",
			match: true,
		},
		{
			name:  "weekday with date with trailing whitespaces",
			input: " 	Piątek 23.10 	",
			match: true,
		},
		{
			name:  "weekday without date",
			input: "Piątek ",
			match: false,
		},
		{
			name:  "weekday with malformed date",
			input: "Piątek 9.10",
			match: false,
		},
		{
			name:  "only weekday",
			input: "Piątek",
			match: false,
		},
		{
			name:  "only date",
			input: "23.10",
			match: false,
		},
		{
			name:  "couple of dates",
			input: "Jadłospis od 09.10 do 13.10",
			match: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// when
			res := matchDay(test.input)

			// then
			assert.Equal(t, test.match, res)
		})
	}
}
