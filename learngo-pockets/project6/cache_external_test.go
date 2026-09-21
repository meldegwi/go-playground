package cache_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"cache"

	"github.com/alecthomas/assert/v2"
)

func TestCache_Parallel_Goroutines(t *testing.T) {
	c := cache.New[int, string](time.Millisecond * 100)

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
	c := cache.New[int, string](time.Millisecond * 100)

	t.Run("write six", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "six")
	})

	t.Run("write boo", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "boo")
	})
}

func TestCache_TTL(t *testing.T) {
	t.Parallel()

	c := cache.New[string, string](time.Millisecond * 100)

	c.Upsert("Egyptian", "Red")

	got, found := c.Read("Egyptian")
	assert.True(t, found)
	assert.Equal(t, "Red", got)

	time.Sleep(time.Millisecond * 200)

	got, found = c.Read("Egyptian")
	assert.False(t, found)
	assert.Equal(t, "", got)
}
