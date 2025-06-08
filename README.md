go-vimeworld is a client for VimeWorld written in go.

[![Go Reference](https://pkg.go.dev/badge/github.com/EpicStep/go-vimeworld/vimeworld.svg)](https://pkg.go.dev/github.com/EpicStep/go-vimeworld/vimeworld)
[![Lint](https://github.com/EpicStep/go-vimeworld/actions/workflows/lint.yml/badge.svg)](https://github.com/EpicStep/go-vimeworld/actions/workflows/lint.yml)
[![Test](https://github.com/EpicStep/go-vimeworld/actions/workflows/test.yml/badge.svg)](https://github.com/EpicStep/go-vimeworld/actions/workflows/test.yml)
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
