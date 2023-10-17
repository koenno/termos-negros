package filter

import (
	"testing"
	"time"

	"github.com/koenno/termos-negros/domain"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnFilteredElements(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		input    domain.Menu
		expected domain.Menu
	}{
		{
			name: "some elements are after the given date",
			date: newTime("2023-10-17 10:54:31"),
			input: []domain.DayMenu{
				newDayMenu("2023-10-17 10:54:32"),
				newDayMenu("2023-10-18 00:00:00"),
				newDayMenu("2023-10-16 23:51:43"),
			},
			expected: []domain.DayMenu{
				newDayMenu("2023-10-17 10:54:32"),
				newDayMenu("2023-10-18 00:00:00"),
			},
		},
		{
			name: "all elements are after the given date",
			date: newTime("2023-10-17 10:54:31"),
			input: []domain.DayMenu{
				newDayMenu("2023-10-17 10:54:32"),
				newDayMenu("2023-10-18 00:00:00"),
				newDayMenu("2023-10-19 23:51:43"),
			},
			expected: []domain.DayMenu{
				newDayMenu("2023-10-17 10:54:32"),
				newDayMenu("2023-10-18 00:00:00"),
				newDayMenu("2023-10-19 23:51:43"),
			},
		},
		{
			name: "all elements are before the given date",
			date: newTime("2023-10-19 23:51:44"),
			input: []domain.DayMenu{
				newDayMenu("2023-10-17 10:54:32"),
				newDayMenu("2023-10-18 00:00:00"),
				newDayMenu("2023-10-19 23:51:43"),
			},
			expected: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// when
			filtered := Since(test.date, test.input)

			// then
			assert.ElementsMatch(t, test.expected, filtered)
		})
	}
}

func newTime(date string) time.Time {
	t, _ := time.Parse(time.DateTime, date)
	return t
}

func newDayMenu(date string) domain.DayMenu {
	return domain.DayMenu{
		Date:  newTime(date),
		Meals: []domain.Meal{},
	}
}
