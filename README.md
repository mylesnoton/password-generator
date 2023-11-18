# Password Generator

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mylesnoton/password-generator/blob/master/LICENSE)

A very simple tool to generate passwords of a specified length and complexity, can be used from the command line or imported as a package. Uses Go's build in random number generator, although hasn't been tested for true randomness as yet.

## Installation

To use the password-generator binary from the command line, run:

```bash
make install
```

If you are using it as a package, run:

```bash
go get github.com/mylesnoton/password-generator/generator
```

## Command Line Usage

To generate a 32 character password with a combination of lower and uppercase letters, numbers and special chars, use:

```bash
password-generator --length 32
```

To generate a 32 character password with a combination of lower and uppercase letters and numbers without an special chars, use:

```bash
password-generator --length 32 --no-special
```

## Package Usage

You can import the generator package into your own project and implement it like this.

```go
package main

import (
    "github.com/mylesnoton/password-generator/generator"
    "fmt"
)

func main() {
    fmt.Println(generator.NewPassword(32))
}
```

## Testing

Tests are included, using the built in Go testing package, to run the tests, execute the following command:

`go test github.com/mylesnoton/password-generator/...`

## Improvements

* Test the randomness of the passwords properly
* Make it more useful
* Allow pronouncable words
* Provide an API maybe

## License

Password Generator is distributed using the MIT License, see LICENSE