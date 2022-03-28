package fifo_test

import (
	"fmt"

	"github.com/floatdrop/fifo"
)

func ExampleLRU() {
	queue := fifo.New[string, int](256)

	queue.Push("Hello", 5)

	if e := queue.Contains("Hello"); e != nil {
		fmt.Println(*e)
		// Output: 5
	}
}
