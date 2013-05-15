package interfaces

import (
	"bufio"
	"encoding/csv"
	. "github.com/ghthor/gospec"
	"os"
	"strconv"
	"time"
	"unicode/utf8"
)

const FILENAME = "games.csv"

func CSVExportSpec(c Context) {
	c.Specify("Persists a single game to csv.", func() {

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

		StoreGamesInCSVFile(FILENAME, []*Game{expectedGame})

		file, _ := os.Open(FILENAME)

		fileReader := bufio.NewReader(file)
		csvReader := csv.NewReader(fileReader)
		delimiter, _ := utf8.DecodeRuneInString(";")
		csvReader.Comma = delimiter
		games, _ := csvReader.ReadAll()
		file.Close()

		removeFile := true

		c.Expect(games[0][0], Equals, "2013-02-28")
		c.Expect(games[0][1], Equals, strconv.Itoa(expectedGame.MatchDay))
		c.Expect(games[0][2], Equals, strconv.Itoa(expectedGame.Position))
		c.Expect(games[0][3], Equals, expectedGame.HomeTeam)
		c.Expect(games[0][4], Equals, expectedGame.HomePlayer1)
		c.Expect(games[0][5], Equals, expectedGame.HomePlayer2)
		c.Expect(games[0][6], Equals, strconv.Itoa(expectedGame.HomeScore))
		c.Expect(games[0][7], Equals, strconv.Itoa(expectedGame.GuestScore))
		c.Expect(games[0][8], Equals, expectedGame.GuestPlayer1)
		c.Expect(games[0][9], Equals, expectedGame.GuestPlayer2)
		c.Expect(games[0][10], Equals, expectedGame.GuestTeam)

		if removeFile {
			os.Remove(FILENAME)
		}
	})
	c.Specify("Persists a double game to csv.", func() {

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

		StoreGamesInCSVFile(FILENAME, []*Game{expectedDoubleGame})

		file, _ := os.Open(FILENAME)

		fileReader := bufio.NewReader(file)
		csvReader := csv.NewReader(fileReader)
		delimiter, _ := utf8.DecodeRuneInString(";")
		csvReader.Comma = delimiter
		games, _ := csvReader.ReadAll()
		file.Close()

		removeFile := true

		c.Expect(games[0][0], Equals, "2013-02-27")
		c.Expect(games[0][1], Equals, strconv.Itoa(expectedDoubleGame.MatchDay))
		c.Expect(games[0][2], Equals, strconv.Itoa(expectedDoubleGame.Position))
		c.Expect(games[0][3], Equals, expectedDoubleGame.HomeTeam)
		c.Expect(games[0][4], Equals, expectedDoubleGame.HomePlayer1)
		c.Expect(games[0][5], Equals, expectedDoubleGame.HomePlayer2)
		c.Expect(games[0][6], Equals, strconv.Itoa(expectedDoubleGame.HomeScore))
		c.Expect(games[0][7], Equals, strconv.Itoa(expectedDoubleGame.GuestScore))
		c.Expect(games[0][8], Equals, expectedDoubleGame.GuestPlayer1)
		c.Expect(games[0][9], Equals, expectedDoubleGame.GuestPlayer2)
		c.Expect(games[0][10], Equals, expectedDoubleGame.GuestTeam)

		if removeFile {
			os.Remove(FILENAME)
		}
	})
}
