package verify

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// MockReader is an interface that offers most of os.File reading operations.
type MockReader interface {
	io.Reader
	io.ReaderAt
	io.Seeker
}

// MockROFile returns a bytes.Buffer backed mock File that implements io.Reader,
// io.Seeker, io.ReaderAt interfaces.
func MockROFile(content string) MockReader {
	return strings.NewReader(content)
}

// FileExists checks if provided path exists.
func FileExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.New(path + " does not exist")
	}
	if err != nil {
		return fmt.Errorf("cannot stat '%s': %s", path, err)
	}
	return nil
}

// FileDoesNotExist checks if provided path does not exist.
func FileDoesNotExist(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("cannot stat '%s': %s", path, err)
	}
	return errors.New(path + " does exist (but should not)")
}

// DirHasContent checks that the given dir contains the provided tree. Tree
// should be relative to root.
func DirHasContent(root string, tree []string) error {
	ls, err := lsDir(root)
	if err != nil {
		return err
	}

	return EqualSliceWithoutOrder(ls, tree)
}

// DirIsEmpty checks that a dir is empty.
func DirIsEmpty(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Readdir(1); err != io.EOF {
		return errors.New(path + " is not empty (but should be)")
	}
	return nil
}

// TestFolder represents a temporary folder where testing can happen
type TestFolder struct {
	Root string
}

// MustNewTestFolder creates a temporary folder to host a test folder for tb.
// If creation fails, tb is terminated by calling tb.Fatal.
func MustNewTestFolder(tb testing.TB) *TestFolder {
	return &TestFolder{
		Root: tb.TempDir(),
	}
}

// Populate populates a temporary testing folders with the given tree The
// provided list of files should be in the correct order where folders are
// listed before the files they contain (if any). Any provided path is relative
// to the temporary testing folder root. Any path provided without any extension
// is interpreted as being a folder.
func (tmp *TestFolder) Populate(tree []string) error {
	for _, f := range tree {
		path := tmp.Fullpath(f)
		if filepath.Ext(path) == "" {
			if err := os.MkdirAll(path, 0777); err != nil {
				return fmt.Errorf("fail to create temporary folder (path %s): %v", f, err)
			}
		} else {
			if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
				return fmt.Errorf("fail to create file %s: %v", path, err)
			}

			if _, err := os.Create(path); err != nil {
				return fmt.Errorf("fail to create temporary file (path %s): %v", f, err)
			}
		}
	}

	return nil
}

// List returns the list of files and folders contained in the temporary test
// folder. Returned tree is made of relative path to testing folder's root. Order
// is the lexical order (as it uses filepath.Walk under the hood).
func (tmp *TestFolder) List() ([]string, error) {
	ls, err := lsDir(tmp.Root)
	if err != nil {
		return nil, fmt.Errorf("fail to list temporary test folder: %v", err)
	}

	return ls, nil
}

// Glob returns the list of files and folders contained in the temporary test
// folder that match the given pattern. Returned tree is made of relative path
// to testing folder's root. Order is the lexical order (as it uses
// filepath.Walk under the hood).
func (tmp *TestFolder) Glob(pattern string) (tree []string, err error) {
	err = filepath.Walk(tmp.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == tmp.Root {
			return nil
		}

		relpath, _ := filepath.Rel(tmp.Root, path) //Sure that we are not going to fail here

		match, err := filepath.Match(pattern, relpath)
		if err != nil {
			return err
		}
		if !match {
			return nil
		}

		tree = append(tree, relpath)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("fail to list temporary test folder (path %s): %v", tmp.Root, err)
	}

	return
}

// ListWithExt returns the list of files contained in the temporary test folder
// whose extension match the given pattern.  Returned tree is made of relative
// path to testing folder's root. Order is the lexical order (as it uses
// filepath.Walk under the hood).
func (tmp *TestFolder) ListWithExt(ext string) (tree []string, err error) {
	err = filepath.Walk(tmp.Root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == tmp.Root {
			return nil
		}

		if filepath.Ext(path) != ext {
			return nil
		}

		relpath, _ := filepath.Rel(tmp.Root, path) //Sure that we are not going to fail here
		tree = append(tree, relpath)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("fail to list temporary test folder (path %s): %v", tmp.Root, err)
	}

	return
}

// Fullpath returns the complete path to the given relative path from the test
// folder root
func (tmp *TestFolder) Fullpath(relpath string) string {
	return filepath.Join(tmp.Root, filepath.Clean("/"+relpath))
}

// ShouldHaveFile check if provided path exists in the test folder. The path can be
// relative to the test folder's root.
func (tmp *TestFolder) ShouldHaveFile(relpath string) error {
	path := tmp.Fullpath(relpath)
	return FileExists(path)
}

// ShouldNotHaveFile checks that the given files are not in the test folder
func (tmp *TestFolder) ShouldNotHaveFile(wanted string) error {
	path := tmp.Fullpath(wanted)
	return FileDoesNotExist(path)
}

// ShouldHaveContent verifies that the temporary test folder content corresponds
// to the wanted tree.
func (tmp *TestFolder) ShouldHaveContent(wanted []string) error {
	return DirHasContent(tmp.Root, wanted)
}

// ShouldBeEmpty verifies that the temporary test folder is empty.
func (tmp *TestFolder) ShouldBeEmpty() error {
	return DirIsEmpty(tmp.Root)
}

func lsDir(root string) (tree []string, err error) {
	err = filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == root {
			return nil
		}
		relpath, _ := filepath.Rel(root, path) //Sure that we are not going to fail here
        if fi.IsDir() {
            relpath = relpath+string(os.PathSeparator)
        }

		tree = append(tree, relpath)
		return nil
	})

	return
}
