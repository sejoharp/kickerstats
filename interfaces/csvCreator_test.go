package interfaces

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"bufio"
	"encoding/csv"
	"os"
	"testing"
	"time"
	"unicode/utf8"
	"strconv"
)

const FILENAME = "games.csv"

func TestStoreGamesInCSVFile(t *testing.T) {
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

	file, err := os.Open(FILENAME)
	helper.HandleFatalError("error creating csv file: ", err)
	fileReader := bufio.NewReader(file)
	csvReader := csv.NewReader(fileReader)
	delimiter, _ := utf8.DecodeRuneInString(";")
	csvReader.Comma = delimiter
	games, err := csvReader.ReadAll()
	helper.HandleFatalError("error reading file: ", err)
	err = file.Close()
	helper.HandleFatalError("error closing file: ", err)

	removeFile := true
	if games[1][0] != "2013-02-27 20:00:00 +0100 CET" {
		t.Errorf("Parsing MatchDate failed. result:%s", games[1][0])
		removeFile = false
	}
	if games[1][1] != strconv.Itoa(expectedDoubleGame.MatchDay) {
		t.Errorf("Parsing MatchDate failed. result:   ", games[1][1])
		removeFile = false
	}
	if games[1][2] != "3" {
		t.Errorf("Parsing position failed. result: ", games[1][2])
		removeFile = false
	}
	if games[1][3] != "Tingeltangel FC St. Pauli" {
		t.Errorf("Parsing hometeam failed. result: ", games[1][3])
		removeFile = false
	}
	if games[1][4] != "Fischer, Harro" {
		t.Errorf("Parsing homeplayer1 failed. result: ", games[1][4])
		removeFile = false
	}
	if games[1][5] != "Kränz, Ludwig" {
		t.Errorf("Parsing homeplayer2 failed. result: ", games[1][5])
		removeFile = false
	}
	if games[1][6] != "4" {
		t.Errorf("Parsing homescore failed. result: ", games[1][6])
		removeFile = false
	}
	if games[1][7] != "5" {
		t.Errorf("Parsing guestscore failed. result: ", games[1][7])
		removeFile = false
	}
	if games[1][8] != "Zierott, Ulli" {
		t.Errorf("Parsing guestplayer1 failed. result: ", games[1][8])
		removeFile = false
	}
	if games[1][9] != "Hojas, René" {
		t.Errorf("Parsing guestplayer2 failed. result: ", games[1][9])
		removeFile = false
	}
	if games[1][10] != "Hamburg Privateers 08" {
		t.Errorf("Parsing guestteam failed. result: ", games[1][10])
		removeFile = false
	}
	if removeFile {
		os.Remove(FILENAME)
	}
}
