package main

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"bitbucket.org/joscha/kickerstats/interfaces"
	"time"
)

func main() {
	startTime := time.Now()
	helper.LogInfo("start")
	allgames := make([]*interfaces.Game, 0)
	seasonRaw := interfaces.DownloadHTMLData("http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe")
	seasonDoc := interfaces.GenerateDocument(seasonRaw)
	seasonIds := interfaces.FindSeasons(seasonDoc)
	for _, seasonId := range seasonIds {
		helper.LogInfo("seasonid :" + seasonId)
		seasonRaw := interfaces.DownloadSeason(seasonId)
		seasonDoc := interfaces.GenerateDocument(seasonRaw)
		ligaLinks := interfaces.FindLigaLinks(seasonDoc)
		for _, ligaLink := range ligaLinks {
			helper.LogInfo("ligaLink: " + ligaLink)
			matchesRaw := interfaces.DownloadHTMLData(ligaLink)
			matchesDoc := interfaces.GenerateDocument(matchesRaw)
			matchLinks := interfaces.FindMatchLinks(matchesDoc)
			for _, matchLink := range matchLinks {
				helper.LogInfo("matchLink: " + matchLink)
				gamesRaw := interfaces.DownloadHTMLData(matchLink)
				gamesDoc := interfaces.GenerateDocument(gamesRaw)
				games := interfaces.ParseGames(gamesDoc)
				allgames = append(allgames, games...)
			}
		}
	}
	
	interfaces.StoreGamesInCSVFile("allGames.csv", allgames)
	duration := time.Since(startTime)
	helper.LogInfo("duration: " + duration.String())
}
