# radius [![GoDoc](https://godoc.org/github.com/bobrovde/radius?status.svg)](https://godoc.org/github.com/bobrovde/radius)

a Go (golang) [RADIUS](https://tools.ietf.org/html/rfc2865) client and server implementation

## Installation

    go get -u github.com/bobrovde/radius

## Client example

```go
package main

import (
	"context"
	"fmt"

	"github.com/bobrovde/radius"
	. "github.com/bobrovde/radius/rfc2865"
)

func main() {
	packet := radius.New(radius.CodeAccessRequest, []byte(`secret`))
	UserName_SetString(packet, "tim")
	UserPassword_SetString(packet, "12345")
	response, err := radius.Exchange(context.Background(), packet, "localhost:1812")
	if err != nil {
		panic(err)
	}

	if response.Code == radius.CodeAccessAccept {
		fmt.Println("Accepted")
	} else {
		fmt.Println("Denied")
	}
}
```

## RADIUS Dictionaries

Included in this package is the command line program `radius-dict-gen`. It can be installed with:

    go get -u github.com/bobrovde/radius/cmd/radius-dict-gen

This program will generate helper functions and types for reading and manipulating RADIUS attributes in a packet. It is recommended that generated code be used for any RADIUS dictionary you would like to consume.

Included in this repository are sub-packages of generated helpers for commonly used RADIUS attributes, including [`rfc2865`](https://godoc.org/github.com/bobrovde/radius/rfc2865) and [`rfc2866`](https://godoc.org/github.com/bobrovde/radius/rfc2866).

## License

MPL 2.0
