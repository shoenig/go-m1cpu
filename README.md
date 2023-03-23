# m1cpu

The `m1cpu` package is a library for inspecting Apple Silicon CPUs in Go.

Use the `m1cpu` library for looking up the CPU frequency for Apple M1 and M2 CPUs.

# Install

```shell
go get github.com/shoenig/m1cpu@latest
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

# License

Open source under the [MPL](LICENSE)
