# m1cpu

[![Go Reference](https://pkg.go.dev/badge/github.com/shoenig/go-m1cpu.svg)](https://pkg.go.dev/github.com/shoenig/go-m1cpu)
[![MPL License](https://img.shields.io/github/license/shoenig/go-m1cpu?color=g&style=flat-square)](https://github.com/shoenig/go-m1cpu/blob/main/LICENSE)
[![Run CI Tests](https://github.com/shoenig/go-m1cpu/actions/workflows/ci.yaml/badge.svg)](https://github.com/shoenig/go-m1cpu/actions/workflows/ci.yaml)

The `go-m1cpu` module is a library for inspecting Apple Silicon CPUs in Go.

Use the `m1cpu` library for looking up the CPU frequency for Apple M1 and M2 CPUs.

# Install

```shell
go get github.com/shoenig/go-m1cpu@latest
```

# Example

```go
package main

import (
    "fmt"

    "github.com/shoenig/go-m1cpu"
)

func main() {
    fmt.Println("pCore", m1cpu.PCoreGHz())
    fmt.Println("eCore", m1cpu.ECoreGHz())
}
```

# CGO

This package requires the use of CGO. Extracting the CPU properties is done via
Apple's IOKit framework, which is accessible only through system C libraries.

# License

Open source under the [MPL](LICENSE)
