# go-piper

[![Build Status](https://travis-ci.org/jasonfriedland/go-piper.svg?branch=master)](https://travis-ci.org/jasonfriedland/go-piper)

A small utility to read from stdin if it was piped in.

## Usage

Lint, test:

    make lint test

Example:

    import "github.com/jasonfriedland/go-piper"
    input, err := piper.Read()
