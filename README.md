# Password Generator

[![CircleCI](https://circleci.com/gh/mylesnoton/passgen.svg?style=svg&circle-token=01ef058c08b109581863046351ca41e38eb8a07e)](https://circleci.com/gh/mylesnoton/passgen)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mylesnoton/passgen/blob/master/LICENSE)

A very simple tool to generate passwords of a specified length and complexity, can be used from the command line or imported as a package. Uses Go's build in random number generator, although hasn't been tested for true randomness as yet.

## Installation

To use the passgen binary from the command line, run:

```bash
$ make install
```

If you are using it as a package, run:

```bash
$ go get github.com/mylesnoton/passgen/generator
```

## Command Line Usage

To generate a 32 character password with a combination of lower and uppercase letters, numbers and special chars, use:

```bash
$ passgen --length 32
```

To generate a 32 character password with a combination of lower and uppercase letters and numbers without an special chars, use:

```bash
$ passgen --length 32 --no-special
```

## Package Usage

You can import the generator package into your own project and implement it like this.

```
package main

import (
    "github.com/mylesnoton/passgen/generator"
    "fmt"
)

func main() {
    fmt.Println(generator.NewPassword(32))
}
```

## Tests

Tests are included, using the built in Go testing package, to run the tests, execute the following command:

`go test github.com/mylesnoton/passgen/...`

## Improvements

* Test the randomness of the passwords properly
* Make it more useful
* Allow pronouncable words
* Provide an API maybe

## License
Password Generator is distributed using the MIT License, see LICENSE