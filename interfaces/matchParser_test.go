package interfaces

import (
	//"fmt"
	"testing"
)

func TestFindMatchLinksAmount(t *testing.T) {
	doc := loadDoc("begegnungen.html")
	expectedMatchLinkCount := 14

	matchLinks := FindMatchLinks(doc)
	if expectedMatchLinkCount != len(matchLinks) {
		t.Errorf("False amount of match links. expected: %d, result: %d", expectedMatchLinkCount, len(matchLinks))
	}
}

func TestFindMatchLinksCheckFirstLink(t *testing.T) {
	doc := loadDoc("begegnungen.html")
	expectedMatchLink := "http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe?task=begegnung_spielplan&veranstaltungid=64&id=3815"

	matchLinks := FindMatchLinks(doc)
	if expectedMatchLink != matchLinks[0] {
		t.Errorf("Parsing first match link failed. expected: %s, result: %s", expectedMatchLink, matchLinks[0])
		t.Errorf(expectedMatchLink)
		t.Errorf(matchLinks[0])
	}
}

func TestFindLigaLinks(t *testing.T) {
	doc := loadDoc("uebersicht.html")

	FindLigaLinks(doc)
}
