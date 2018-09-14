# exopulse times package
Golang package **times** contains common date/time operations.

[![CircleCI](https://circleci.com/gh/exopulse/times.svg?style=svg)](https://circleci.com/gh/exopulse/times)
[![Build Status](https://travis-ci.org/exopulse/times.svg?branch=master)](https://travis-ci.org/exopulse/times)
[![GitHub license](https://img.shields.io/github/license/exopulse/times.svg)](https://github.com/exopulse/times/blob/master/LICENSE)

# Overview

This package contains common date/time operations.

## Features

### Date/time parse with format auto-detection. 

# Using times package

## Installing package

Use go get to install the latest version of the library.

    $ go get github.com/exopulse/times
 
Include times in your application.
```go
import "github.com/exopulse/times"
```

## Use Parse to parse input string as time. This method tries to detect input format.
```go
t, err := times.Parse("2006-01-02T15:04:05-0700")
```

# About the project

## Contributors

* [exopulse](https://github.com/exopulse)

## License

Times package is released under the MIT license. See
[LICENSE](https://github.com/exopulse/times/blob/master/LICENSE)
