package filter

import (
	"time"

	"github.com/koenno/termos-negros/domain"
)

func Since(t time.Time, menu domain.Menu) domain.Menu {
	var filtered domain.Menu
	for _, m := range menu {
		if m.Date.Compare(t) >= 0 {
			filtered = append(filtered, m)
		}
	}
	return filtered
}
