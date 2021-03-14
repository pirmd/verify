package verify

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/sanity-io/litter"

	"github.com/pirmd/text/diff"
	"github.com/pirmd/text/table"
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
			h = []diff.Highlighter{diff.WithSoftTabs, diff.WithoutMissingContent}
		}
		if *showdiffcolor {
			h = []diff.Highlighter{diff.WithSoftTabs, diff.WithColor}
		}
		if *showdiffNP {
			h = append([]diff.Highlighter{diff.WithNonPrintable}, h...)
		}

		dL, dR, dT, _ := delta.PrettyPrint(h...)
		return table.New().AddCol(dR, dT, dL).SetHeader("Got", "", "Want").String()
	}

	return fmt.Sprintf("Got:\n%v\n\nWant :\n%v", g, w)
}

func stringify(v interface{}) string {
	//TODO(pirmd): github.com/sanity-io/litter use strconv.Quote on each
	//string, which is not what I usually look for so I bypass it.
	//Find something more sensible (forking litter and using options?)
	s := litter.Sdump(v)
	if u, err := strconv.Unquote(s); err == nil {
		return u
	}
	return s
}
