# Changelog
All notable changes to this project will be documented in this file.

Format is based on [Keep a Changelog] (https://keepachangelog.com/en/1.0.0/).
Versionning adheres to [Semantic Versioning] (https://semver.org/spec/v2.0.0.html)

## [Unrelease]
### Modified
- Break previous principles for verify functions. Verify functions are now
  outputting errors in case of check fail and leave the end user the choice on
  how to handle the situation in its unit testing code (for example t.Errorf or
  t.Fatalf...)
- Switch to github.com/pirmd/text v0.4.0 that offers better diff algorithm and
  color diff formatting.
- Use github.com/sanity-io/litter to provide readily usable pretty formatting
  of compared values. 


## [0.3.0] - 2019-11-11
### Added
- Add extended error message definition in MatchGolden
- Add function to list all files matching a pattern from a TestField
- Add ListWithExt function to TestField
- Add helpers to capture os.Stdout and compare it to either a string or a
  golden file
- Add basic mock http server that alters http.DefaultTransport and replace it
  to serve locally stored http.Response
### Modified
- Change comparison functions to always expect 'got' and 'want' arguments in
  the same order


## [0.2.0] - 2019-08-11
### Added
- Add a flag to displays differences in colors
### Modified
- Correct a BUG when displaying diff in test results
- Modified slightly the formatting of error messages in case of tests' errors 

## [0.1.0] - 2019-05-11
### Added
- basic comparison with optional diff output between 'expected' and 'got' result
- support of golden files and there management (updating)
- basic function to play with temporary testing files
