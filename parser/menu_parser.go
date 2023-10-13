package parser

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/koenno/termos-negros/domain"
	"golang.org/x/exp/slog"
)

var (
	ErrParse = errors.New("failed to parse the document")
)

type MenuParser struct {
}

func (p *MenuParser) Parse(r io.Reader) ([]domain.DayMenu, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("failed to create document: %v", err)
	}

	var menu []domain.DayMenu
	doc.Find(".card-header").Each(func(_ int, div *goquery.Selection) {
		text := div.Text()
		if !matchDay(text) {
			return
		}

		dayHeader := strings.TrimSpace(div.Text())
		split := strings.Split(dayHeader, " ")
		if len(split) < 2 {
			slog.Error("malformed day header", "header", dayHeader)
			return
		}
		dateWithYear := fmt.Sprintf("%s.%d", split[1], time.Now().Year())
		date, err := time.Parse("02.01.2006", dateWithYear)
		if err != nil {
			slog.Error("failed to parse date", "date", split[2])
			return
		}
		dayMenu := domain.DayMenu{
			Date: date,
		}
		i := 0
		meal := domain.Meal{}
		div.Siblings().Find("p").Each(func(_ int, p *goquery.Selection) {
			text = p.Text()
			if i%2 == 0 {
				meal.Name = text
			} else {
				meal.Ingredients = text
				dayMenu.Meals = append(dayMenu.Meals, meal)
			}
			i++
		})
		menu = append(menu, dayMenu)
	})

	return menu, nil
}
