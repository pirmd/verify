package verify

import (
	"flag"
	"fmt"

	"github.com/sanity-io/litter"

	"github.com/pirmd/text"
	"github.com/pirmd/text/diff"
)

var (
	showdiff      = flag.Bool("test.diff", false, "show differences between result and expected values")
	showdiffcolor = flag.Bool("test.diff-color", false, "show differences using colors between result and expected values")
	showdiffNP    = flag.Bool("test.diff-np", false, "show differences between result and expected values materializing non printable chars")
)

func msgWithDiff(got, want interface{}) string {
	g, w := stringify(got), stringify(want)

	if *showdiff || *showdiffcolor || *showdiffNP {
		delta := diff.Patience(w, g, diff.ByLines, diff.ByRunes)

		h := []diff.Highlighter{}
		if *showdiff {
			h = []diff.Highlighter{diff.WithSoftTabs}
		}
		if *showdiffcolor {
			h = []diff.Highlighter{diff.WithSoftTabs, diff.WithColor}
		}
		if *showdiffNP {
			h = append([]diff.Highlighter{diff.WithNonPrintable}, h...)
		}

		dL, dR, dT := delta.PrettyPrint(h...)
		//XXX:return text.NewTable().Col(dR, dT, dL).Captions("Got", "", "Want").Draw()
		return text.NewTable().Rows([]string{dL, dT, dR}).Captions("Got", "", "Want").Draw()
	}

	return fmt.Sprintf("Got:\n%v\n\nWant :\n%v", g, w)
}

func stringify(v interface{}) string {
	return litter.Sdump(v)
}
