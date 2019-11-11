# VERIFY
[![GoDoc](https://godoc.org/github.com/pirmd/verify?status.svg)](https://godoc.org/github.com/pirmd/verify)&nbsp; 
[![Go Report Card](https://goreportcard.com/badge/github.com/pirmd/verify)](https://goreportcard.com/report/github.com/pirmd/verify)&nbsp;

`verify` package is a collection of simple functions that I find handy to use
for building my unit testing.

It is not a full testing framework nor expected to cover all testing needs
under the sky, just some tools I happen to use on the top of the standard go
testing module.

# INSTALLATION
Everything should work fine using go standard commands (`build`, `get`,
`install`...).

# USAGE
Running `go doc github.com/pirmd/verify` should give you helpful guidelines on
avail bales features.

With the package set-up, additional go test flags are offered:
    . `-test.golden-update`: updates the golden files with the test result.
    . `-test.mockhttp-update`: updates files served through the mock-http transport.
    . `-test.diff`: show differences between the test result and expected values.
    . `-test.colordiff`: show differences in color between the test result and the
       expected values.

# CONTRIBUTION
If you feel like to contribute, just follow github guidelines on
[forking](https://help.github.com/articles/fork-a-repo/) then [send a pull
request](https://help.github.com/articles/creating-a-pull-request/)

[modeline]: # ( vim: set fenc=utf-8 spell spl=en: )
