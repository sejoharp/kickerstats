package interfaces

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"fmt"
	"github.com/puerkitobio/goquery"
	"strconv"
	"strings"
	"time"
)

func ParseGames(doc *goquery.Document) (games []*Game) {
	rawGames := doc.Find("div#Content > table.contentpaneopen:nth-child(6) > tbody")
	rawGames = rawGames.Find(".sectiontableentry1, .sectiontableentry2")
	hasImages := hasImages(rawGames)
	homeTeam, guestTeam := parseTeams(doc)
	matchDate := parseMatchDate(doc)
	matchDay := parseMatchDay(doc)
	rawGames.Each(func(i int, selection *goquery.Selection) {
		if isValidGame(selection) {
			game := &Game{}
			game.Double = isDouble(selection)
			game.Position = parseGamePosition(selection)
			game.HomeTeam = homeTeam
			game.GuestTeam = guestTeam
			game.MatchDate = matchDate
			game.MatchDay = matchDay
			addPlayers(game, selection)
			if hasImages {
				scoreString := selection.Children().Eq(3).Text()
				game.HomeScore, game.GuestScore = parseScores(scoreString)
			} else {
				scoreString := selection.Children().Eq(2).Text()
				game.HomeScore, game.GuestScore = parseScores(scoreString)
			}
			games = append(games, game)
		}
	})
	return
}

func isValidGame(selection *goquery.Selection) bool {
	tds := selection.Children().Length()
	return tds == 6 || tds == 4
}

func parseMatchDate(doc *goquery.Document) time.Time {
	rawData := doc.Find("#Content table tbody > tr > td").First().Text()
	dateChunk := strings.Split(rawData, ",")[1]
	return parseDate(strings.TrimSpace(dateChunk))
}

func parseMatchDay(doc *goquery.Document) int {
	rawData := doc.Find("#Content table tbody > tr > td").First().Text()
	dateChunks := strings.Split(rawData, ",")
	matchDayString := strings.Split(dateChunks[2], ".")[0]
	matchDay, err := strconv.Atoi(strings.TrimSpace(matchDayString))
	helper.HandleFatalError("parsing matchday failed:", err)
	return matchDay
}

func parseTeams(doc *goquery.Document) (homeTeam string, guestTeam string) {
	fmt.Print("")
	teams := doc.Find("table.contentpaneopen").Eq(1).Find("tbody > tr > td > table > tbody h2")
	homeTeam = teams.First().Text()
	guestTeam = teams.Last().Text()
	return
}

func parseGamePosition(selection *goquery.Selection) (position int) {
	positionString := strings.TrimSpace(selection.Children().First().Text())
	var err error
	position, err = strconv.Atoi(strings.TrimSpace(positionString))
	helper.HandleFatalError("parsing game position failed:", err)
	return
}

func parseScores(scores string) (homeScore int, guestScore int) {
	var err error
	scoreChunks := strings.Split(scores, ":")
	homeScore, err = strconv.Atoi(strings.TrimSpace(scoreChunks[0]))
	helper.HandleFatalError("parsing home score failed:", err)
	guestScore, err = strconv.Atoi(strings.TrimSpace(scoreChunks[1]))
	helper.HandleFatalError("parsing guest score failed:", err)
	return
}

func isDouble(selection *goquery.Selection) bool {
	return selection.Find("td a").Length() == 4
}

func addPlayers(game *Game, selection *goquery.Selection) {
	if game.Double {
		game.HomePlayer1 = selection.Find("td a").Eq(0).Text()
		game.HomePlayer2 = selection.Find("td a").Eq(1).Text()
		game.GuestPlayer1 = selection.Find("td a").Eq(2).Text()
		game.GuestPlayer2 = selection.Find("td a").Eq(3).Text()
	} else {
		game.HomePlayer1 = selection.Find("td a").Eq(0).Text()
		game.GuestPlayer1 = selection.Find("td a").Eq(1).Text()
	}
}

func hasImages(rawGames *goquery.Selection) bool {
	return rawGames.First().Children().Length() == 6
}
