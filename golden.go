package verify

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	//GoldenDir is the path where to find golden files
	GoldenDir = "./testdata"

	// updateGolden golden files, when set, updates the content of the golden
	// files from th etest result.
	updateGolden = flag.Bool("test.golden-update", false, "update golden file with test result")
)

// MatchGolden compares a test result to the content of a 'golden' file
// If 'update' command flag is used, update the 'golden' file
func MatchGolden(name string, got string) error {
	goldenPath := filepath.Join(GoldenDir, name+".golden")

	if *updateGolden {
		if err := updateGoldenFiles(goldenPath, []byte(got)); err != nil {
			return err
		}
	}

	want, err := readGolden(goldenPath)
	if err != nil {
		return fmt.Errorf("cannot read golden file %s: %s.\nTest output is:\n%s", goldenPath, err, got)
	}

	return Equal(got, string(want))
}

func readGolden(path string) ([]byte, error) {
	want, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
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
