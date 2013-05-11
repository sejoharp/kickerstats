package main

import (
	"bitbucket.org/joscha/kickerstats/interfaces"
	"github.com/ghthor/gospec"
	"testing"
)

func TestAllInterfacesSpecs(t *testing.T) {
	r := gospec.NewRunner()

	r.AddSpec(interfaces.MatchParserSpec)
	r.AddSpec(interfaces.GameParserSpec)
	r.AddSpec(interfaces.CSVExportSpec)
	//r.AddSpec(interfaces.LigatoolReaderSpec)

	gospec.MainGoTest(r, t)
}
