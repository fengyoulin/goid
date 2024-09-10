package goid

import (
	"sort"
	"sync"
	"testing"
)

func TestID(t *testing.T) {
	limit := 50
	ids := make([]int, limit)
	ch := make(chan uint64, limit)
	wg := sync.WaitGroup{}
	wg.Add(limit)
	for limit > 0 {
		limit--
		go func() {
			ch <- ID()
			wg.Done()
		}()
	}
	wg.Wait()
	close(ch)
	for i := 0; i < len(ids); i++ {
		ids[i] = int(<-ch)
	}
	sort.Ints(ids)
	for i := 1; i < len(ids); i++ {
		if ids[i-1] == ids[i] {
			t.Errorf("duplicate: %d", ids[i])
		}
	}
}
