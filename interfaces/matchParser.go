package interfaces

import (
	//"bitbucket.org/joscha/hpfeed/helper"
	"fmt"
	"github.com/puerkitobio/goquery"
	"strings"
)

func FindMatchLinks(doc *goquery.Document) (matchLinks []string) {

	rawMatchLinks := doc.Find("div#Content > table.contentpaneopen:nth-child(7) > tbody")
	rawMatchLinks = rawMatchLinks.Find(".sectiontableentry1, .sectiontableentry2")
	rawMatchLinks.Each(func(i int, selection *goquery.Selection) {
		alreadyPlayed := selection.Find("a").Length() == 2
		scoreConfirmed := selection.Find("td:nth-child(5) small").Text() != "unbestätigt"
		if alreadyPlayed && scoreConfirmed {
			link, _ := selection.Find("a[href]").Attr("href")
			matchLinks = append(matchLinks, "http://www.kickern-hamburg.de"+strings.TrimSpace(link))
		}
	})
	return
}

func FindLigaLinks(doc *goquery.Document) (ligaLinks []string) {
	rawLigaLinks := doc.Find("div#Content > table > tbody > tr > td > a.readon")
	fmt.Print(rawLigaLinks.Text())
	return
}
