package interfaces

import (
	"bufio"
	"encoding/csv"
	. "github.com/ghthor/gospec"
	"github.com/ghthor/gospec"
	"os"
	"strconv"
	"time"
	"unicode/utf8"
)

const FILENAME = "games.csv"

func CSVExportSpec(c gospec.Context) {
	c.Specify("Persists two games to csv.", func() {

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

		StoreGamesInCSVFile(FILENAME, []*Game{expectedGame, expectedDoubleGame})

		file, _ := os.Open(FILENAME)

		fileReader := bufio.NewReader(file)
		csvReader := csv.NewReader(fileReader)
		delimiter, _ := utf8.DecodeRuneInString(";")
		csvReader.Comma = delimiter
		games, _ := csvReader.ReadAll()
		file.Close()

		removeFile := true

		c.Expect(games[1][0], Equals, "2013-02-27 20:00:00 +0100 CET")
		c.Expect(games[1][1], Equals, strconv.Itoa(expectedDoubleGame.MatchDay))
		c.Expect(games[1][2], Equals, strconv.Itoa(expectedDoubleGame.Position))
		c.Expect(games[1][3], Equals, expectedDoubleGame.HomeTeam)
		c.Expect(games[1][4], Equals, expectedDoubleGame.HomePlayer1)
		c.Expect(games[1][5], Equals, expectedDoubleGame.HomePlayer2)
		c.Expect(games[1][6], Equals, strconv.Itoa(expectedDoubleGame.HomeScore))
		c.Expect(games[1][7], Equals, strconv.Itoa(expectedDoubleGame.GuestScore))
		c.Expect(games[1][8], Equals, expectedDoubleGame.GuestPlayer1)
		c.Expect(games[1][9], Equals, expectedDoubleGame.GuestPlayer2)
		c.Expect(games[1][10], Equals, expectedDoubleGame.GuestTeam)

		if removeFile {
			os.Remove(FILENAME)
		}
	})
}
