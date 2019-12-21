# goid #

Get current goroutine's ID, which is maintained by the golang runtime and can be used as an unique identifier for the goroutine.

Example:
```go
package main

import (
	"fmt"
	"github.com/fengyoulin/goid"
	"sync"
)

const limit = 50

func main() {
	wg := sync.WaitGroup{}
	wg.Add(limit)

	for i := 0; i < limit; i++ {
		go func(idx int) {
			fmt.Printf("%d: %d\n", idx, goid.ID())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```
