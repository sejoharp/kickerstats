package interfaces

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"bytes"
	"exp/html"
	"github.com/puerkitobio/goquery"
	"time"
	"fmt"
)

const RAW_DATE_FORMAT = "02.01.2006 15:04"

type Player struct {
	playerId string
	name     string
	team     string
}

type Game struct {
	HomePlayer1  string
	HomePlayer2  string
	HomeTeam     string
	HomeScore    int
	GuestPlayer1 string
	GuestPlayer2 string
	GuestTeam    string
	GuestScore   int
	Position     int
	MatchDate    time.Time
	MatchDay     int
	Double       bool
}

func (this *Game) Equal(other *Game) bool {
	return this.HomePlayer1 == other.HomePlayer1 &&
		this.HomePlayer2 == other.HomePlayer2 &&
		this.HomeTeam == other.HomeTeam &&
		this.HomeScore == other.HomeScore &&
		this.GuestPlayer1 == other.GuestPlayer1 &&
		this.GuestPlayer2 == other.GuestPlayer2 &&
		this.GuestTeam == other.GuestTeam &&
		this.GuestScore == other.GuestScore &&
		this.Position == other.Position &&
		this.MatchDate == other.MatchDate &&
		this.MatchDay == other.MatchDay &&
		this.Double == other.Double
}

func (this *Game) EqualDebug(other *Game) (equals bool) {
	equals = true
	if this.HomePlayer1 != other.HomePlayer1 {
		equals = false
		fmt.Println("HomePlayer1")
	}
	if this.HomePlayer2 != other.HomePlayer2 {
		equals = false
		fmt.Println("HomePlayer2")
	}
	if this.HomeTeam != other.HomeTeam {
		equals = false
		fmt.Println("HomeTeam")
	}
	if this.HomeScore != other.HomeScore {
		equals = false
		fmt.Println("HomeScore")
	}
	if this.GuestPlayer1 != other.GuestPlayer1 {
		equals = false
		fmt.Println("GuestPlayer1")
	}
	if this.GuestPlayer2 != other.GuestPlayer2 {
		equals = false
		fmt.Println("GuestPlayer2")
	}
	if this.GuestTeam != other.GuestTeam {
		equals = false
		fmt.Println("GuestTeam")
	}
	if this.GuestScore != other.GuestScore {
		equals = false
		fmt.Println("GuestScore")
	}
	if this.Position != other.Position {
		equals = false
		fmt.Println("Position")
	}
	if this.MatchDate != other.MatchDate {
		equals = false
		fmt.Println("MatchDate")
	}
	if this.MatchDay != other.MatchDay {
		equals = false
		fmt.Println("MatchDay")
	}
	if this.Double != other.Double {
		equals = false
		fmt.Println("Double")
	}
	return
}

func (this *Game) Print() {
	fmt.Println(this.MatchDate,
		this.MatchDay,
		this.Position,
		this.Double,
		this.HomePlayer1,
		this.HomePlayer2,
		this.HomeTeam,
		this.HomeScore,
		" vs. ",
		this.GuestScore,
		this.GuestPlayer1,
		this.GuestPlayer2,
		this.GuestTeam)
}

func PrintAllGames(games []*Game) {
	for _, game := range games {
		game.Print()
	}
}

func toUtf8(iso8859_1_buf []byte) string {
	buf := make([]rune, len(iso8859_1_buf))
	for i, b := range iso8859_1_buf {
		buf[i] = rune(b)
	}
	return string(buf)
}

func parseDate(rawDate string) time.Time {
	date, _ := time.Parse(RAW_DATE_FORMAT, rawDate)
	return overrideLocation(date)
}

func overrideLocation(t time.Time) time.Time {
	y, m, d := t.Date()
	H, M, S := t.Clock()
	return time.Date(y, m, d, H, M, S, 0, time.Local)
}

func GenerateDocument(rawData []byte) *goquery.Document {
	utf8String := toUtf8(rawData)
	utf8byteArray := []byte(utf8String)
	node, err := html.Parse(bytes.NewReader(utf8byteArray))
	helper.HandleFatalError("document generation failed:", err)
	return goquery.NewDocumentFromNode(node)
}
