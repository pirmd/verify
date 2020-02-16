# VERIFY
[![GoDoc](https://godoc.org/github.com/pirmd/verify?status.svg)](https://godoc.org/github.com/pirmd/verify)&nbsp; 
[![Go Report Card](https://goreportcard.com/badge/github.com/pirmd/verify)](https://goreportcard.com/report/github.com/pirmd/verify)&nbsp;

`verify` package is a collection of simple functions that I find handy to use
for building my unit testing.

# INSTALLATION
Everything should work fine using go standard commands (`build`, `get`,
`install`...).

# USAGE
Running `go doc github.com/pirmd/verify` should give you helpful guidelines on
availables features.

With the package set-up, additional go test flags are offered:
    . `-test.golden-update`: updates the golden files with the test result.
    . `-test.mockhttp-update`: updates files served through the mock-http transport.
    . `-test.diff`: show differences between the test result and expected values.
    . `-test.diff-color`: show differences in color between the test result and the
       expected values.
    . `-test.diff-np`: show differences between result and expected values
      materializing non printable chars.

# CONTRIBUTION
If you feel like to contribute, just follow github guidelines on
[forking](https://help.github.com/articles/fork-a-repo/) then [send a pull
request](https://help.github.com/articles/creating-a-pull-request/)

[modeline]: # ( vim: set fenc=utf-8 spell spl=en: )
