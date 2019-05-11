package verify

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type testField struct {
	tb   testing.TB
	Root string
}

//NewTestField creates a temporary folder to host a test field for tb.
//If a temporary test field already exists for tb, it fails.
func NewTestField(tb testing.TB) *testField {
	path, err := ioutil.TempDir("", tb.Name())
	if err != nil {
		tb.Fatalf("Cannot create temporary test field: %v", err)
	}

	return &testField{tb, path}
}

//Clean removes the temporary test fields created for tb if it exists
func (td *testField) Clean() {
	if err := os.RemoveAll(td.Root); err != nil {
		td.tb.Fatalf("Fail to remove temporary test field: %v", err)
	}
}

//Populate populates a temporary testing fields with the given tree
//The provided list of files should be in the correct order where folders are
//listed before the files they contain (if any). Any provided path is relative
//to the temporary testing field root. Any path provided without any extension
//is interpreted as being a folder.
func (td *testField) Populate(tree []string) {
	for _, f := range tree {
		path := td.Fullpath(f)
		if filepath.Ext(path) == "" {
			if err := os.MkdirAll(path, 0777); err != nil {
				td.tb.Fatalf("Fail to create temporary folder (path %s): %v", f, err)
			}
		} else {
			if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
				td.tb.Fatalf("Fail to create file %s: %v", path, err)
			}

			if _, err := os.Create(path); err != nil {
				td.tb.Fatalf("Fail to create temporary file (path %s): %v", f, err)
			}
		}
	}
}

//List returns the list of files and folders contained in th etemporary test
//field. Returned tree is made of relative path to testing field's root.
//Order is fiwed (lexical order as it uses filepath.Walk under the hood)
func (td *testField) List() (tree []string) {
	err := filepath.Walk(td.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == td.Root {
			return nil
		}
		relpath, _ := filepath.Rel(td.Root, path) //Sure that we are not going to fail here
		tree = append(tree, relpath)
		return nil
	})

	if err != nil {
		td.tb.Fatalf("Fail to list temporary test field: %v", err)
	}

	return
}

//Fullname returns the complete path to the given relative path from the
//test field root
func (td *testField) Fullpath(relpath string) string {
	return filepath.Join(td.Root, filepath.Clean("/"+relpath))
}

//Exists check if provided path exists in the test field. The path can be
//relative to the test field's root.
func (td *testField) Exists(relpath string) bool {
	path := td.Fullpath(relpath)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		td.tb.Errorf("Cannot stat '%s': %s", path, err)
		return false
	}
	return true
}

//ShouldHaveContent verifies that the temporary test field content corresponds to the
//wanted tree.
func (td *testField) ShouldHaveContent(wanted []string, message string) {
	EqualSliceWithoutOrder(td.tb, td.List(), wanted, message)
}

//ShouldHaveFile verifies that the wanted file/folder is in the test field
func (td *testField) ShouldHaveFile(wanted string, message string) {
	if !td.Exists(wanted) {
		td.tb.Logf("Test field should contain %s", wanted)
		td.tb.Error(message)
	}
}

//ShouldNotHaveFiles checks that the given files are not in the test field
func (td *testField) ShouldNotHaveFile(wanted string, message string) {
	if td.Exists(wanted) {
		td.tb.Logf("Test field should not contain %s", wanted)
		td.tb.Error(message)
	}
}

//IOReader returns a simple io.Reader, usefull to mock files
func IOReader(content string) io.Reader {
	return bytes.NewBufferString(content)
}
