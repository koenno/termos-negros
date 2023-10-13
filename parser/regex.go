package parser

import (
	"fmt"
	"regexp"
)

const (
	weekDays = "Poniedziałek|Wtorek|Środa|Czwartek|Piątek"
)

func matchDay(str string) bool {
	regex := fmt.Sprintf("^\\s*%s [0-9][0-9]\\.[0-9][0-9]\\s*$", weekDays)
	r := regexp.MustCompile(regex)
	return r.MatchString(str)
}
