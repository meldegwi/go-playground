package cache_test

import (
	"fmt"
	"sync"
	"testing"

	"cache"
)

func TestCache_Parallel_Goroutines(t *testing.T) {
	c := cache.New[int, string]()

	const parallelTasks = 10

	wg := sync.WaitGroup{}

	for i := 0; i <= parallelTasks; i++ {
		wg.Go(func() {
			c.Upsert(4, fmt.Sprint(i))
		})
	}

	wg.Wait()
}

func TestCache_Parallel(t *testing.T) {
	c := cache.New[int, string]()

	t.Run("write six", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "six")
	})

	t.Run("write boo", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "boo")
	})
}
