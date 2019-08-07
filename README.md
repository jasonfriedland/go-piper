# go-piper

[![Build Status](https://travis-ci.org/jasonfriedland/go-piper.svg?branch=master)](https://travis-ci.org/jasonfriedland/go-piper)
[![codecov](https://codecov.io/gh/jasonfriedland/go-piper/branch/master/graph/badge.svg)](https://codecov.io/gh/jasonfriedland/go-piper)

A small utility for use in CLI programs, to read piped-in input if it was
supplied.

## Usage

Lint, test:

    make lint test

Example:

    import "github.com/jasonfriedland/go-piper"

    input, err := piper.Read()
