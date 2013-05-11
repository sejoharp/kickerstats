package interfaces

import (
	. "github.com/ghthor/gospec"
)

func LigatoolReaderSpec(c Context) {
	c.Specify("Downloads full html from the 4th season.", func() {
		rawdata := DownloadSeason("4")
		doc := GenerateDocument(rawdata)
		c.Expect(doc.Length(), Equals, 1)
	})
	c.Specify("Download html data.", func() {
		rawdata := DownloadHTMLData("http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe?task=begegnung_spielplan&veranstaltungid=8&id=2")
		doc := GenerateDocument(rawdata)
		c.Expect(doc.Length(), Equals, 1)
	})
}
