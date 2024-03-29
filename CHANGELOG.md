# Changelog
## [0.7.0] - 2022-03-20
- Simplify files-related testing by taking benefit of new
  testing.T.TempDir()

## [0.6.0] - 2021-03-14
- Update dependencies to github.com/pirmd/text v0.6.0
- Get rid of external dependency (github.com/sanity-io/litter).

## [0.5.3] - 2020-12-01
### Modified
- Get rid of external dependency (github.com/mattetti/filebuffer).

## [0.5.2] - 2020-11-29
### Modified
- Extend MockROFile to support io.ReaderAt and io.Seeker interfaces.

## [0.5.1] - 2020-07-12
### Added
- Add function to verify that a folder is empty

## [0.5.0] - 2020-05-16
### Added
- Simple wrapper around testing.TB to allow plumbing testing log facilities to
  an existing logger.

## [O.4.2] - 2020-05-08 
### Modified
- FIX diff without color to show empty right text when an insertion is detected.
- Improve support/feedback to user in case golden files do not exist.

## [0.4.1] - 2020-02-17
### Modified
- FIX an unwanted behavior following introduction of
  github.com/sanity-io/litter as stringifier where all strings where quoted
  using strconv.Quote. It is not the behavior I expect in most cases, bypass it
  as a short term workaround.

## [0.4.0] - 2020-02-16
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
