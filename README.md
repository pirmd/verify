# INTRODUCTION
`verify` package is a collection of simple functions that I find handy to use
for building my unit testing.

It is not a full testing framework nor expected to cover all testing needs
under the sky, just some tools I happen to use on the top of the standard go
testing module.

# INSTALLATION
Everything should work fine using go standard commands (`build`, `get`,
`install`...).

# USAGE
Running `godoc` should give you helpful guideines on availbales features.

With the package set-up, two additional go tst flags are offered:
    - `-test.update`: updates the golden files with the test result
    - `-test.diff`  : show differences between the test result and the expected
      test result

# CONTRIBUTION
If you feel like to contribute, just follow github guidelines on
[forking](https://help.github.com/articles/fork-a-repo/) then [send a pull
request](https://help.github.com/articles/creating-a-pull-request/)

[modeline]: # ( vim: set fenc=utf-8 spell spl=en: )
