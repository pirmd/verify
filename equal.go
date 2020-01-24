package verify

import (
	"errors"
	"reflect"
	"sort"
)

// Equal verifies that 'got' is equal to 'want'
func Equal(got, want interface{}) error {
	if !reflect.DeepEqual(got, want) {
		return errors.New(msgWithDiff(got, want))
	}

	return nil
}

// EqualSliceWithoutOrder verifies that two slices of strings are equal whatever
// the order of their content is
func EqualSliceWithoutOrder(got, want []string) error {
	sort.Strings(got)
	sort.Strings(want)
	return Equal(got, want)
}

// Panic verifies that the given func will panic when run
func Panic(f func()) (err error) {
	defer func() {
		if r := recover(); r == nil {
			err = errors.New("func did not panic as expected")
		}
	}()
	f()
	return
}
