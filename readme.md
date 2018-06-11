# Go string set

A simple library implementing a string set data type.

## Installation

```
go get -u github.com/gtrafimenkov/stringset
```

## Usage

```
ss := stringset.New("foo", "bar")
ss.Add("foo")
fmt.Println(ss)
```

## Tests

```
go test --cover
```

## License

MIT-0
