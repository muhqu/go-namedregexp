Go NamedRegexp package
======================

[![Build Status](https://travis-ci.org/muhqu/go-namedregexp.png?branch=master)](https://travis-ci.org/muhqu/go-namedregexp)

Adds a few methods to aid with working with named group submatches.

```go
func (*NamedRegexp) FindNamedStringSubmatch(string) map[string]string
func FindNamedStringSubmatch(*regexp.Regexp, string) map[string]string
func (*NamedRegexp) FindNamedStringSubmatchIndex(string) map[string][]int
func FindNamedStringSubmatchIndex(*regexp.Regexp, string) map[string][]int
```

Usage Example
-------------
```go
import (
  namedregexp "github.com/muhqu/go-namedregexp"
)

var myExp = namedregexp.MustCompile(`^(?P<firstname>\w+)\s+(?P<lastname>\w+)$`)

func main() {
  submatches := myExp.FindNamedStringSubmatch("John Doe")
  submatches["firstname"] // John
  submatches["lastname"] // Doe

  submatchesIndex := myExp.FindNamedStringSubmatchIndex("John Doe")
  submatchesIndex["firstname"] // []int{0, 4}
  submatchesIndex["lastname"] // []int{5, 8}
}
```


