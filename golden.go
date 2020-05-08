package verify

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	// updateGolden golden files
	updateGolden = flag.Bool("test.golden-update", false, "update golden file with test result")
	// where to find golden files
	goldenDir = flag.String("test.goldendir", "./testdata", "path to folder hosting golden files")
)

// MatchGolden compares a test result to the content of a 'golden' file
// If 'update' command flag is used, update the 'golden' file
func MatchGolden(name string, got string) error {
	goldenPath := filepath.Join(*goldenDir, name+".golden")

	if *updateGolden {
		if err := updateGoldenFiles(goldenPath, []byte(got)); err != nil {
			return err
		}
	}

	want, err := readGolden(goldenPath)
	if err != nil {
		return fmt.Errorf("cannot read golden file %s: %s.\nTest output is:\n%s", goldenPath, err, got)
	}

	if len(want) == 0 {
		return fmt.Errorf("no existing or empty golden file.\nTest output is:\n%s", got)
	}

	return Equal(got, string(want))
}

func readGolden(path string) ([]byte, error) {
	want, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, fmt.Errorf("cannot read golden file %s: %v", path, err)
	}
	return want, nil
}

func updateGoldenFiles(path string, actual []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return fmt.Errorf("cannot update golden file %s: %v", path, err)
	}

	if err := ioutil.WriteFile(path, actual, 0666); err != nil {
		return fmt.Errorf("cannot update golden file %s: %v", path, err)
	}

	return nil
}
