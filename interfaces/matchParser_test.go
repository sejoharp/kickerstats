package interfaces

import (
	"github.com/ghthor/gospec"
	. "github.com/ghthor/gospec"
	"strconv"
)

func MatchParserSpec(c gospec.Context) {
	c.Specify("Parser returns all match links.", func() {
		doc := loadDoc("begegnungen.html")

		matchLinks := FindMatchLinks(doc)

		c.Expect(len(matchLinks), Equals, 14)
	})

	c.Specify("The first match link is filled.", func() {
		doc := loadDoc("begegnungen.html")

		matchLinks := FindMatchLinks(doc)

		c.Expect(matchLinks[0], Equals, "http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe?task=begegnung_spielplan&veranstaltungid=64&id=3815")
	})

	c.Specify("Parser returns all ligalinks.", func() {
		doc := loadDoc("uebersicht.html")

		links := FindLigaLinks(doc)

		c.Expect(len(links), Equals, 5)
	})

	c.Specify("Parser returns all seasons.", func() {
		doc := loadDoc("uebersicht.html")

		seasonIds := FindSeasons(doc)

		c.Expect(len(seasonIds), Equals, 5)
	})

	c.Specify("The first seasonID is filled.", func() {
		doc := loadDoc("uebersicht.html")

		seasonIds := FindSeasons(doc)
		seasonId, _ := strconv.Atoi(seasonIds[0])

		c.Expect(seasonId, Equals, 7)
	})
}
