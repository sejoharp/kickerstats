package interfaces

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"github.com/puerkitobio/goquery"
	"strconv"
	"strings"
	"time"
)

func ParseGames(doc *goquery.Document) (games []*Game) {
	rawGames := doc.Find("div#Content > table.contentpaneopen:nth-child(6) > tbody")
	rawGames = rawGames.Find(".sectiontableentry1, .sectiontableentry2")
	if isValidGameList(rawGames) {
		hasImages := hasImages(rawGames)
		homeTeam, guestTeam := parseTeams(doc)
		var matchDate time.Time
		hasMatchDate := hasMatchDate(doc)
		matchDate = parseMatchDate(doc, hasMatchDate)
		matchDay := parseMatchDay(doc, hasMatchDate)
		rawGames.Each(func(i int, selection *goquery.Selection) {
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
		})
	}
	return
}

func isValidGameList(selection *goquery.Selection) bool {
	tdCount := selection.First().Children().Length()
	return tdCount == 6 || tdCount == 4
}

func hasMatchDate(doc *goquery.Document) bool {
	rawData := doc.Find("#Content table tbody > tr > td").First().Text()
	dateChunks := strings.Split(rawData, ",")
	return len(dateChunks) == 3
}
func parseMatchDate(doc *goquery.Document, hasMatchDate bool) time.Time {
	if hasMatchDate == false {
		return time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	}
	rawData := doc.Find("#Content table tbody > tr > td").First().Text()
	dateChunks := strings.Split(rawData, ",")
	dateChunk := dateChunks[1]
	return parseDate(strings.TrimSpace(dateChunk))
}

func parseMatchDay(doc *goquery.Document, hasMatchDate bool) int {
	rawData := doc.Find("#Content table tbody > tr > td").First().Text()
	dateChunks := strings.Split(rawData, ",")
	var matchDayString string
	if hasMatchDate {
		matchDayString = strings.Split(dateChunks[2], ".")[0]
	} else {
		matchDayString = strings.Split(dateChunks[1], ".")[0]
	}
	matchDay, err := strconv.Atoi(strings.TrimSpace(matchDayString))
	helper.HandleFatalError("parsing matchday failed:", err)
	return matchDay
}

func parseTeams(doc *goquery.Document) (homeTeam string, guestTeam string) {
	teams := doc.Find("table.contentpaneopen").Eq(1).Find("tbody > tr > td > table > tbody h2")
	homeTeam = removeTeamDescriptons(teams.First().Text())
	guestTeam = removeTeamDescriptons(teams.Last().Text())
	return
}

func removeTeamDescriptons(team string) string {
	return strings.TrimSpace(strings.Split(team, "(")[0])
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
	players := selection.Find("td a")
	if game.Double {
		game.HomePlayer1 = parseName(players, 0)
		game.HomePlayer2 = parseName(players, 1)
		game.GuestPlayer1 = parseName(players, 2)
		game.GuestPlayer2 = parseName(players, 3)
	} else {
		game.HomePlayer1 = parseName(players, 0)
		game.GuestPlayer1 = parseName(players, 1)

	}
}

func parseName(players *goquery.Selection, position int) string {
	if players.Length()-1 >= position {
		return players.Eq(position).Text()
	}
	return ""
}

func hasImages(rawGames *goquery.Selection) bool {
	return rawGames.First().Children().Length() == 6
}
