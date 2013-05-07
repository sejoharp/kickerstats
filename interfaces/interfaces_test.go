// Copyright © 2009-2011 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package interfaces

import (
	"exp/html"
	"fmt"
	"github.com/ghthor/gospec"
	"github.com/puerkitobio/goquery"
	"os"
	"testing"
)

// You will need to list every spec in a TestXxx method like this,
// so that gotest can be used to run the specs. Later GoSpec might
// get its own command line tool similar to gotest, but for now this
// is the way to go. This shouldn't require too much typing, because
// there will be typically only one top-level spec per class/feature.

func TestAllInterfacesSpecs(t *testing.T) {
	r := gospec.NewRunner()

	// List all specs here
	r.AddSpec(MatchParserSpec)
	r.AddSpec(GameParserSpec)
	r.AddSpec(CSVExportSpec)
	//r.AddSpec(LigatoolReaderSpec)

	// Run GoSpec and report any errors to gotest's `testing.T` instance
	gospec.MainGoTest(r, t)
}

func loadDoc(page string) *goquery.Document {
	if f, e := os.Open(fmt.Sprintf("../testdata/%s", page)); e != nil {
		panic(e.Error())
	} else {
		defer f.Close()
		if node, e := html.Parse(f); e != nil {
			panic(e.Error())
		} else {
			return goquery.NewDocumentFromNode(node)
		}
	}
	return nil
}
