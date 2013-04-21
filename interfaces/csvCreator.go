package interfaces

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"unicode/utf8"
)

func StoreGamesInCSVFile(fileName string, games []*Game) {
	file, err := os.Create(fileName)
	helper.HandleFatalError("error creating csv file: ", err)
	fileWriter := bufio.NewWriter(file)
	delimiter, _ := utf8.DecodeRuneInString(";")
	csvWriter := csv.NewWriter(fileWriter)
	csvWriter.Comma = delimiter
	for _, game := range games {
		record := []string{
			game.MatchDate.String(),
			strconv.Itoa(game.MatchDay),
			strconv.Itoa(game.Position),
			game.HomeTeam,
			game.HomePlayer1,
			game.HomePlayer2,
			strconv.Itoa(game.HomeScore),
			strconv.Itoa(game.GuestScore),
			game.GuestPlayer1,
			game.GuestPlayer2,
			game.GuestTeam}
		csvWriter.Write(record)
	}
	csvWriter.Flush()
	err = fileWriter.Flush()
	helper.HandleFatalError("error flushing file writer: ", err)
	err = file.Close()
	helper.HandleFatalError("error closing file: ", err)
}
