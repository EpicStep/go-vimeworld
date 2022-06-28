go-vimeworld is a client for VimeWorld written in go.

[![Go Reference](https://pkg.go.dev/badge/github.com/EpicStep/go-vimeworld/vimeworld.svg)](https://pkg.go.dev/github.com/EpicStep/go-vimeworld/vimeworld)
[![tests](https://github.com/EpicStep/go-vimeworld/actions/workflows/tests.yml/badge.svg)](https://github.com/EpicStep/go-vimeworld/actions/workflows/tests.yml)
[![x](https://github.com/EpicStep/go-vimeworld/actions/workflows/ci.yml/badge.svg)](https://github.com/EpicStep/go-vimeworld/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/EpicStep/go-vimeworld/branch/master/graph/badge.svg?token=IY1T4ZMYZ8)](https://codecov.io/gh/EpicStep/go-vimeworld)
[![Go Report Card](https://goreportcard.com/badge/github.com/EpicStep/go-vimeworld)](https://goreportcard.com/report/github.com/EpicStep/go-vimeworld)

----

## Installation
```bash
go get github.com/EpicStep/go-vimeworld
```

## Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/EpicStep/go-vimeworld/vimeworld"
)

func main() {
	c, err := vimeworld.NewClient(vimeworld.Options{})
	if err != nil {
		panic(err)
	}

	user, err := c.GetUsersByNames(context.Background(), "EpicStep")
	if err != nil {
		panic(err)
	}

	if len(user) == 0 {
		panic("user not found")
	}

	fmt.Println(user[0])
}
```

## License
MIT