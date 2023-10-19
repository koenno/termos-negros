package filter

import (
	"time"

	"github.com/koenno/termos-negros/domain"
)

func Since(t time.Time, menu domain.Menu) domain.Menu {
	var filtered domain.Menu
	t, _ = time.Parse(time.DateOnly, t.Format(time.DateOnly))
	for _, m := range menu {
		if t.Compare(m.Date) <= 0 {
			filtered = append(filtered, m)
		}
	}
	return filtered
}
