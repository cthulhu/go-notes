[![Build Status](https://travis-ci.org/cthulhu/go-notes.svg?branch=master)](https://travis-ci.org/cthulhu/go-notes)  [![Goreport](https://goreportcard.com/badge/github.com/cthulhu/go-notes)](https://goreportcard.com/badge/github.com/cthulhu/go-notes) [![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/cthulhu/go-notes/master/LICENSE)

# go-notes
CLI tool similar to "rake notes" from Ruby On Rails

# Installation

    go get -u github.com/cthulhu/go-notes

# Usage

    Usage: go-notes [flags] <Go file or directory> ...

    Without options generates all the note types. Default are:

    // FIXME    - call to fix something
    // OPTIMIZE - call for a refactoring
    // TODO     - future plans

    Options:
      -f - FIXME annotations
      -o - OPTIMIZE annotations
      -t - TODO annotations
