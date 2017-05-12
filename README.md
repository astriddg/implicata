# implicata

[![Build Status](https://travis-ci.org/danbondd/implicata.svg?branch=master)](https://travis-ci.org/danbondd/implicata)

A simple application that monitors and logs events from a HTML form to catch possible fraudulent activity.

## Requirements

- Go 1.8.x

## Usage

- Download source code `go get github.com/danbondd/implicata`
- Navigate to project `cd $GOPATH/src/github.com/danbondd/implicata`
- Run `make run`
- Visit `http://localhost:8080` in a web browser.
- Interact with form and monitor `stdout`

## Test

- Run `make test`
