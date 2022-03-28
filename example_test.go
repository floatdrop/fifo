package fifo_test

import (
	"fmt"

	"github.com/floatdrop/fifo"
)

func ExampleFIFO() {
	queue := fifo.New[string, int](256)

	queue.Push("Hello", 5)

	if e := queue.Get("Hello"); e != nil {
		fmt.Println(*e)
		// Output: 5
	}
}
