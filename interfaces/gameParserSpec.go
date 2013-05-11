package interfaces

import (
	"exp/html"
	"fmt"
	. "github.com/ghthor/gospec"
	"github.com/puerkitobio/goquery"
	"os"
	"time"
)

func GameParserSpec(c Context) {
	c.Specify("Parsing match returns correct match date.", func() {
		doc := loadDoc("begegnung.html")

		resultDate := parseMatchDate(doc, true)

		c.Expect(time.Date(2013, 2, 27, 20, 0, 0, 0, time.Local), Equals, resultDate)
	})
	c.Specify("Parsing match returns correct match day.", func() {
		doc := loadDoc("begegnung.html")

		resultDay := parseMatchDay(doc, true)

		c.Expect(resultDay, Equals, 1)
	})
	c.Specify("Parsing opponents returns both teams.", func() {
		doc := loadDoc("begegnung.html")

		homeTeam, guestTeam := parseTeams(doc)
		c.Expect(homeTeam, Equals, "Tingeltangel FC St. Pauli")
		c.Expect(guestTeam, Equals, "Hamburg Privateers 08")
	})
	c.Specify("Parsing games returns a full game.", func() {
		doc := loadDoc("begegnung.html")

		games := ParseGames(doc)

		c.Expect(games[0].HomeTeam, Equals, "Tingeltangel FC St. Pauli")
		c.Expect(games[0].GuestTeam, Equals, "Hamburg Privateers 08")
		c.Expect(games[0].HomePlayer1, Equals, "Technau, Jerome")
		c.Expect(games[0].GuestPlayer1, Equals, "Hojas, René")
		c.Expect(games[0].HomeScore, Equals, 4)
		c.Expect(games[0].GuestScore, Equals, 7)
		c.Expect(games[0].Position, Equals, 1)
		c.Expect(games[0].MatchDay, Equals, 1)
		c.Expect(games[0].MatchDate, Equals, time.Date(2013, 2, 27, 20, 0, 0, 0, time.Local))
		c.Expect(games[0].Double, Equals, false)
	})
	c.Specify("Parsing games returns a full double game.", func() {
		doc := loadDoc("begegnung.html")

		games := ParseGames(doc)

		c.Expect(games[2].HomeTeam, Equals, "Tingeltangel FC St. Pauli")
		c.Expect(games[2].GuestTeam, Equals, "Hamburg Privateers 08")
		c.Expect(games[2].HomePlayer1, Equals, "Fischer, Harro")
		c.Expect(games[2].HomePlayer2, Equals, "Kränz, Ludwig")
		c.Expect(games[2].GuestPlayer1, Equals, "Zierott, Ulli")
		c.Expect(games[2].GuestPlayer2, Equals, "Hojas, René")
		c.Expect(games[2].HomeScore, Equals, 4)
		c.Expect(games[2].GuestScore, Equals, 5)
		c.Expect(games[2].Position, Equals, 3)
		c.Expect(games[2].MatchDay, Equals, 1)
		c.Expect(games[2].MatchDate, Equals, time.Date(2013, 2, 27, 20, 0, 0, 0, time.Local))
		c.Expect(games[2].Double, Equals, true)
	})
	c.Specify("Parser returns all games.", func() {
		doc := loadDoc("begegnung.html")

		games := ParseGames(doc)

		c.Expect(len(games), Equals, 16)
	})
	c.Specify("Parsing games with player images returns all games.", func() {
		doc := loadDoc("begegnung_bild.html")

		games := ParseGames(doc)

		c.Expect(len(games), Equals, 16)
	})

	c.Specify("Parsing games with player images returns a full game.", func() {
		doc := loadDoc("begegnung_bild.html")

		games := ParseGames(doc)

		c.Expect(games[0].HomeTeam, Equals, "Cim Bom Bom")
		c.Expect(games[0].GuestTeam, Equals, "Die Maschinerie")
		c.Expect(games[0].HomePlayer1, Equals, "Arslan, Mehmet Emin")
		c.Expect(games[0].GuestPlayer1, Equals, "Bai, Minyoung")
		c.Expect(games[0].HomeScore, Equals, 4)
		c.Expect(games[0].GuestScore, Equals, 7)
		c.Expect(games[0].Position, Equals, 1)
		c.Expect(games[0].MatchDay, Equals, 1)
		c.Expect(games[0].MatchDate, Equals, time.Date(2013, 2, 28, 20, 0, 0, 0, time.Local))
		c.Expect(games[0].Double, Equals, false)
	})
	c.Specify("Parsing games with player images returns a full double game.", func() {
		doc := loadDoc("begegnung_bild.html")

		games := ParseGames(doc)

		c.Expect(games[2].HomeTeam, Equals, "Cim Bom Bom")
		c.Expect(games[2].GuestTeam, Equals, "Die Maschinerie")
		c.Expect(games[2].HomePlayer1, Equals, "Günther, Richard")
		c.Expect(games[2].HomePlayer2, Equals, "Eggerstedt, Kai")
		c.Expect(games[2].GuestPlayer1, Equals, "Bai, Minyoung")
		c.Expect(games[2].GuestPlayer2, Equals, "Strecker, Knuth")
		c.Expect(games[2].HomeScore, Equals, 5)
		c.Expect(games[2].GuestScore, Equals, 4)
		c.Expect(games[2].Position, Equals, 3)
		c.Expect(games[2].MatchDay, Equals, 1)
		c.Expect(games[2].MatchDate, Equals, time.Date(2013, 2, 28, 20, 0, 0, 0, time.Local))
		c.Expect(games[2].Double, Equals, true)
	})
	c.Specify("Parsing relegation does not find games", func() {
		doc := loadDoc("relegation.html")

		games := ParseGames(doc)

		c.Expect(len(games), Equals, 0)
	})
	c.Specify("Parsing games with empty player names return all available data.", func() {
		doc := loadDoc("begegnung_no_names.html")

		games := ParseGames(doc)

		c.Expect(games[13].HomeTeam, Equals, "Die Hinkelsteinchen")
		c.Expect(games[13].GuestTeam, Equals, "Kurbelkraft Bergedorf")
		c.Expect(games[13].HomePlayer1, Equals, "")
		c.Expect(games[13].GuestPlayer1, Equals, "")
		c.Expect(games[13].HomeScore, Equals, 7)
		c.Expect(games[13].GuestScore, Equals, 0)
		c.Expect(games[13].Position, Equals, 14)
		c.Expect(games[13].MatchDay, Equals, 1)
		c.Expect(games[13].MatchDate, Equals, time.Date(2013, 2, 28, 20, 0, 0, 0, time.Local))
		c.Expect(games[13].Double, Equals, false)
	})
	c.Specify("Parsing games without dates returns match day", func() {
		doc := loadDoc("begegnung_no_date.html")

		resultDay := parseMatchDay(doc, false)

		c.Expect(resultDay, Equals, 5)
	})
	c.Specify("Parsing a match without a date returns a default date", func() {
		doc := loadDoc("begegnung_no_date.html")

		matchDate := hasMatchDate(doc)

		c.Expect(matchDate, IsFalse)

		games := ParseGames(doc)

		c.Expect(games[0].MatchDate, Equals, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local))
	})
}

func loadDoc(page string) *goquery.Document {
	if f, e := os.Open(fmt.Sprintf("testdata/%s", page)); e != nil {
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
