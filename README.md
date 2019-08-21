[![Actions Status](https://github.com/cthulhu/go-notes/workflows/Go/badge.svg)](https://github.com/cthulhu/go-notes/actions)
  [![Goreport](https://goreportcard.com/badge/github.com/cthulhu/go-notes)](https://goreportcard.com/report/github.com/cthulhu/go-notes) [![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/cthulhu/go-notes/master/LICENSE)

# go-notes
CLI tool similar to "rake notes" from Ruby On Rails

# Installation

    go get -u github.com/cthulhu/go-notes

# Usage

    Usage: go-notes [flags] <Go file or directory> ...

    Without options generates all the note types. Defaults are:

    // FIXME    - call to fix something
    // OPTIMIZE - call for a refactoring
    // TODO     - future plans

    Options:
      -f - FIXME annotations
      -o - OPTIMIZE annotations
      -t - TODO annotations
      -c CUSTOM - custom annotation label
      -format count - output format aggregated counts
      -format list - output format list with files and annotations (default)
