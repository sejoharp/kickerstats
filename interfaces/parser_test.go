package interfaces

import (
	"exp/html"
	"fmt"
	"github.com/puerkitobio/goquery"
	"os"
	"testing"
	"time"
)

func loadDoc(page string) *goquery.Document {
	if f, e := os.Open(fmt.Sprintf("../testdata/%s", page)); e != nil {
		panic(e.Error())
	} else {
		defer f.Close()
		if node, e := html.Parse(f); e != nil {
			panic(e.Error())
		} else {
			return goquery.NewDocumentFromNode(node)
		}
	}
	return nil
}

func TestParseMatchDate(t *testing.T) {
	doc := loadDoc("begegnung.html")
	expectedMatchDate := time.Date(2013, 2, 27, 20, 0, 0, 0, time.Local)

	resultDate := parseMatchDate(doc)

	if expectedMatchDate != resultDate {
		t.Errorf("Parsing match date failed. expected: %s, result: %s", expectedMatchDate, resultDate)
	}
}

func TestParseMatchDay(t *testing.T) {
	doc := loadDoc("begegnung.html")

	resultDay := parseMatchDay(doc)

	if resultDay != 1 {
		t.Errorf("Parsing match day failed. expected: %s, result: %s", 1, resultDay)
	}
}

func TestParseTeams(t *testing.T) {
	expectedHomeTeam := "Tingeltangel FC St. Pauli"
	expectedGuestTeam := "Hamburg Privateers 08"
	doc := loadDoc("begegnung.html")

	homeTeam, guestTeam := parseTeams(doc)

	if homeTeam != expectedHomeTeam {
		t.Errorf("Parsing home team failed. expected: %s, result: %s", expectedHomeTeam, homeTeam)
	}
	if guestTeam != expectedGuestTeam {
		t.Errorf("Parsing home team failed. expected: %s, result: %s", expectedGuestTeam, guestTeam)
	}
}

func TestParseFirstGame(t *testing.T) {
	expectedGame := &Game{
		HomeTeam:     "Tingeltangel FC St. Pauli",
		GuestTeam:    "Hamburg Privateers 08",
		HomePlayer1:  "Technau, Jerome",
		GuestPlayer1: "Hojas, René",
		HomeScore:    4,
		GuestScore:   7,
		Position:     1,
		MatchDay:     1,
		MatchDate:    time.Date(2013, 2, 27, 20, 0, 0, 0, time.Local),
		Double:       false}
	doc := loadDoc("begegnung.html")

	games := ParseGames(doc)

	if expectedGame.Equal(games[0]) == false {
		t.Errorf("Parsing first game failed. expected: ", expectedGame)
		t.Errorf("Parsing first game failed.   result: ", games[0])
	}
}

func TestParseFirstDoubleGame(t *testing.T) {
	expectedDoubleGame := &Game{
		HomeTeam:     "Tingeltangel FC St. Pauli",
		GuestTeam:    "Hamburg Privateers 08",
		HomePlayer1:  "Fischer, Harro",
		HomePlayer2:  "Kränz, Ludwig",
		GuestPlayer1: "Zierott, Ulli",
		GuestPlayer2: "Hojas, René",
		HomeScore:    4,
		GuestScore:   5,
		Position:     3,
		MatchDay:     1,
		MatchDate:    time.Date(2013, 2, 27, 20, 0, 0, 0, time.Local),
		Double:       true}
	doc := loadDoc("begegnung.html")

	games := ParseGames(doc)

	if expectedDoubleGame.Equal(games[2]) == false {
		t.Errorf("Parsing first double game failed. expected: ", expectedDoubleGame)
		t.Errorf("Parsing first double game failed.   result: ", games[2])
	}
}

func TestParseGameAmount(t *testing.T) {
	doc := loadDoc("begegnung.html")
	expectedGameAmount := 16

	games := ParseGames(doc)
	gameAmount := len(games)

	if expectedGameAmount != gameAmount {
		t.Errorf("Parsing home team failed. expected: %i, result: %i", expectedGameAmount, gameAmount)
	}
}

func TestParseGameAmountWithImages(t *testing.T) {
	doc := loadDoc("begegnung_bild.html")
	expectedGameAmount := 16

	games := ParseGames(doc)
	gameAmount := len(games)

	if expectedGameAmount != gameAmount {
		t.Errorf("Parsing home team failed. expected: %i, result: %i", expectedGameAmount, gameAmount)
	}
}

func TestParseFirstGameWithImages(t *testing.T) {
	expectedGame := &Game{
		HomeTeam:     "Cim Bom Bom",
		GuestTeam:    "Die Maschinerie",
		HomePlayer1:  "Arslan, Mehmet Emin",
		GuestPlayer1: "Bai, Minyoung",
		HomeScore:    4,
		GuestScore:   7,
		Position:     1,
		MatchDay:     1,
		MatchDate:    time.Date(2013, 2, 28, 20, 0, 0, 0, time.Local),
		Double:       false}
	doc := loadDoc("begegnung_bild.html")

	games := ParseGames(doc)

	if expectedGame.Equal(games[0]) == false {
		t.Errorf("Parsing first game failed. expected: ", expectedGame)
		t.Errorf("Parsing first game failed.   result: ", games[0])
	}
}

func TestParseFirstDoubleGameWithImages(t *testing.T) {
	expectedDoubleGame := &Game{
		HomeTeam:     "Cim Bom Bom",
		GuestTeam:    "Die Maschinerie",
		HomePlayer1:  "Günther, Richard",
		HomePlayer2:  "Eggerstedt, Kai",
		GuestPlayer1: "Bai, Minyoung",
		GuestPlayer2: "Strecker, Knuth",
		HomeScore:    5,
		GuestScore:   4,
		Position:     3,
		MatchDay:     1,
		MatchDate:    time.Date(2013, 2, 28, 20, 0, 0, 0, time.Local),
		Double:       true}
	doc := loadDoc("begegnung_bild.html")

	games := ParseGames(doc)

	if expectedDoubleGame.Equal(games[2]) == false {
		t.Errorf("Parsing first double game failed. expected: ", expectedDoubleGame)
		t.Errorf("Parsing first double game failed.   result: ", games[2])
	}
}
