package verify

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	// updateGolden golden files
	updateGolden = flag.Bool("test.update-golden", false, "update golden file with test result")
	// where to find golden files
	goldenDir = flag.String("test.goldendir", "./testdata", "path to folder hosting golden files")
)

// MatchGolden compares a test result to the content of a 'golden' file
// If 'update' command flag is used, update the 'golden' file
func MatchGolden(tb testing.TB, got string, message ...string) {
	if *updateGolden {
		updateGoldenFiles(tb, []byte(got))
	}

	expected := readGolden(tb)

	if len(expected) == 0 {
		tb.Fatalf("no existing or empty golden file. Test output is:\n%s", got)
	}

	EqualString(tb, got, string(expected), message...)
}

func goldenPath(tb testing.TB) string {
	return filepath.Join(*goldenDir, tb.Name()+".golden")
}

func readGolden(tb testing.TB) []byte {
	f := goldenPath(tb)

	expected, err := ioutil.ReadFile(f)
	if err != nil {
		tb.Logf("cannot read golden file %s: %v", f, err)
		return []byte{}
	}
	return expected
}

func updateGoldenFiles(tb testing.TB, actual []byte) {
	f := goldenPath(tb)

	tb.Logf("update golden file %s", f)
	if err := ioutil.WriteFile(f, actual, os.ModePerm); err != nil {
		tb.Fatalf("cannot update golden file %s: %v", f, err)
	}
}
