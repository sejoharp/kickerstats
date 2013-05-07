package interfaces

import (
	"fmt"
	"github.com/ghthor/gospec"
	. "github.com/ghthor/gospec"
)

func LigatoolReaderSpec(c gospec.Context) {
	c.Specify("Download the 4th season returns full html.", func() {
		rawdata := DownloadSeason("4")
		doc := GenerateDocument(rawdata)
		fmt.Println(doc.Html())

		c.Expect(doc.Length(), Equals, 1)
	})
	c.Specify("Download html data returns full html.", func() {
		rawdata := DownloadHTMLData("http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe?task=begegnung_spielplan&veranstaltungid=8&id=2")
		doc := GenerateDocument(rawdata)
		fmt.Println(doc.Html())
		c.Expect(doc.Length(), Equals, 1)
	})
}
