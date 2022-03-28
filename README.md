# fifo
[![Go Reference](https://pkg.go.dev/badge/github.com/floatdrop/fifo.svg)](https://pkg.go.dev/github.com/floatdrop/fifo)
[![CI](https://github.com/floatdrop/fifo/actions/workflows/ci.yml/badge.svg)](https://github.com/floatdrop/fifo/actions/workflows/ci.yml)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/floatdrop/fifo)](https://goreportcard.com/report/github.com/floatdrop/fifo)

Thread safe GoLang fixed size FIFO with O(1) `Get`.

## Example

```go
import (
	"fmt"

	"github.com/floatdrop/fifo"
)

func main() {
	cache := fifo.New[string, int](256)

	cache.Push("Hello", 5)

	if e := cache.Get("Hello"); e != nil {
		fmt.Println(*e)
		// Output: 5
	}
}
```
