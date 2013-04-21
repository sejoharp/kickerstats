package interfaces

import (
	"fmt"
	"testing"
)

func TestGetSeason(t *testing.T) {
	// rawdata := GetSeason("4")
	// doc := GenerateDocument(rawdata)
	// fmt.Println(doc.Html())
}

func TestGetHTMLData(t *testing.T) {
	rawdata := GetHTMLData("http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe?task=begegnung_spielplan&veranstaltungid=8&id=2")
	doc := GenerateDocument(rawdata)
	fmt.Println(doc.Html())
}
