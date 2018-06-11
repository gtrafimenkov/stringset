# Go string set

[![GoDoc](https://godoc.org/github.com/gtrafimenkov/stringset?status.svg)](https://godoc.org/github.com/gtrafimenkov/stringset)

A simple library implementing a string set data type.

## Installation

```
go get -u github.com/gtrafimenkov/stringset
```

## Usage

```
package main

import (
	"fmt"

	"github.com/gtrafimenkov/stringset"
)

func main() {
	ss := stringset.New("foo", "bar")
	ss.Add("foo")
	fmt.Println(ss)

	// Output: ["bar", "foo"]
}
```

## Documentation

[godoc.org/github.com/gtrafimenkov/stringset](https://godoc.org/github.com/gtrafimenkov/stringset)

## Tests

```
go test --cover
```

## License

MIT-0
