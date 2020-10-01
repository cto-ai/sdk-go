# Project Structure and Release Process

## Project Structure

The public API of this SDK is defined in the source files in the root
directory. The logic in this code communicates with the daemon using
the structures and functions defined in `internal/daemon`.

## Release Process

The Go SDK does not have a functioning release process; the current
git `master` is always the current release.

Notes:
- The documentation for Go module versioning is not very good and it
  is unclear why the process we attempted was not successful
- Unlike our other languages, there is no way to specify a flexible
  dependency version in a `go.mod` file. This causes problems with the
  smooth updating of the SDK to new versions.
