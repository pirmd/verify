package verify

import (
	"encoding/json"
	"flag"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/pirmd/cli/style/text"
)

var (
	showdiff = flag.Bool("test.diff", false, "show differences between result and expected values")
)

//Equal verifies that 'got' is equal to 'want' and feedback a test error
//message with diff information
func Equal(tb testing.TB, got, want interface{}, message ...string) {
	if !reflect.DeepEqual(got, want) {
		errorf(tb, message...)
		if *showdiff {
			dT, dL, dR := text.Diff.Anything(want, got)
			tb.Errorf("\n%s", text.NewTable().Title("Want", "", "Got").Col(dL, dT, dR))
		} else {
			tb.Errorf("\nWant:%#v\nGot :%#v", want, got)
		}
	}
}

//EqualString verifies that 'got' is equal to 'want' and feedback a test error
//message with a line by line diff between them
func EqualString(tb testing.TB, got, want string, message ...string) {
	if got != want {
		errorf(tb, message...)
		if *showdiff {
			dT, dL, dR := text.Diff.Bylines(want, got)
			tb.Errorf("\n%s", text.NewTable().Title("Want", "", "Got").Col(dL, dT, dR))
		} else {
			tb.Errorf("Want:%s\n\nGot :%s", want, got)
		}
	}
}

//EqualAsJSON verifies that Json encoding of got is equal to want's one.  It is
//a weak comparison mean but can be useful to compare data structures that
//relies on interface{} that can have different type but 'similar' content.
func EqualAsJSON(tb testing.TB, got, want interface{}, message ...string) {
	EqualString(tb, stringify(got), stringify(want), message...)
}

//EqualSliceWithoutOrder verifies that two slices of strings are equal whatever
//the order of their content is
func EqualSliceWithoutOrder(tb testing.TB, got, want []string, message ...string) {
	sort.Strings(got)
	sort.Strings(want)
	EqualString(tb, strings.Join(want, "\n"), strings.Join(got, "\n"), message...)
}

//Panic verifies that the given func will panic when run
func Panic(tb testing.TB, f func(), message ...string) {
	defer func() {
		if r := recover(); r == nil {
			errorf(tb, message...)
		}
	}()
	f()
}

func errorf(tb testing.TB, message ...string) {
	if len(message) == 0 {
		tb.Fail()
	}

	s := make([]interface{}, len(message)-1)
	for i, m := range message[1:] {
		s[i] = m
	}
	tb.Errorf(message[0], s...)
}

func stringify(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}

	return string(b)
}
