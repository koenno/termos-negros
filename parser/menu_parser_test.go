package parser

import (
	"os"
	"testing"
	"time"

	"github.com/koenno/termos-negros/domain"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnParsedDayMenu(t *testing.T) {
	// given
	f, err := os.OpenFile("testdata/test-site.html", os.O_RDONLY, 0777)
	assert.NoError(t, err)
	sut := &MenuParser{}

	// when
	menu, err := sut.Parse(f)

	// then
	assert.NoError(t, err)
	assert.NotEmpty(t, menu)
	assert.Equal(t, 9, menu[0].Date.Day())
	assert.Equal(t, time.Month(10), menu[0].Date.Month())
	assert.ElementsMatch(t, expectedMondayMeals(), menu[0].Meals)
	assert.Equal(t, 10, menu[1].Date.Day())
	assert.Equal(t, time.Month(10), menu[1].Date.Month())
	assert.ElementsMatch(t, expectedTuedayMeals(), menu[1].Meals)
}

func expectedMondayMeals() []domain.Meal {
	return []domain.Meal{
		{
			Name:        "Śniadanie",
			Ingredients: "kiełbasa drobiowa (97%mięsa z piersi kurczaka), ser żółty (7)/ser żółty wegański, pomidorki malinowe, rzodkiewki, ogórki zielone, sałata lodowa, rukola, masło i pieczywo żytnie na zakwasie (1) dla diety pieczywo bezglutenowe, półmisek owoców sezonowych, herbatka owocowa",
		},
		{
			Name: "Obiad",
			Ingredients: `Krupnik z kaszą jęczmienną (1) z młodą marchewką, pietruszką, selerem i ziemniakami (9), (dla wersji bezglutenowej z kaszą jaglaną) na wywarze warzywnym, koperek do posypania :)
Dal z dyni hokkaido i czerwonej soczewicy z cebulką, czosnkiem oraz przyprawami z ryżem basmati oraz ogórki gruntowe do pochrupania :)`,
		},
		{
			Name:        "Podwieczorek",
			Ingredients: "chlebek bananowy z mąką owsianą, płatkami owsianymi i kawałkami czekolady (1,3)",
		},
	}
}

func expectedTuedayMeals() []domain.Meal {
	return []domain.Meal{
		{
			Name:        "Śniadanie",
			Ingredients: "jajka na twardo od kur z wolnego wybiegu (3), twarożek z rzodkiewkami i szczypiorkiem (7)/twarożek wegański (8), ogórki zielone, mix młodych zielonych listków, pomidory malinowe, masło i pieczywo żytnie na zakwasie, dla diety pieczywo bezglutenowe, półmisek owoców sezonowych, herbata owocowa",
		},
		{
			Name: "Obiad",
			Ingredients: `Żurek z ziemniakami (1) (na naturalnym zakwasie żytnim lub zakwasie jaglanym dla diety bezglutenowej) z duszoną cebulką, czosnkiem, majerankiem, liściem laurowym, zielem angielskim, kozieradką i pieprzem ziarnistym
Pieczone udka z kurczaka wolnowybiegowego, ziemniaki z koperkiem i masełkiem oraz marchewka z groszkiem wege: kotlety z marynowanego tofu (1,3,6)`,
		},
		{
			Name:        "Podwieczorek",
			Ingredients: "„skrzynki” owocowe ;)",
		},
	}
}
