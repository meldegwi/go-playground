package cache

import (
	"sync"
	"time"
)

// Cache is a gerneric cache, its key can be any comparale value and the attached,
// value can be anything.
type Cache[K comparable, V any] struct {
	ttl  time.Duration
	mu   sync.Mutex
	data map[K]EntryWithTimeout[V]
}

type EntryWithTimeout[V any] struct {
	value   V
	expires time.Time // after that time the value is expired.
}

// New makes a new instance of the Cache struct.
func New[K comparable, V any](ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:  ttl,
		data: make(map[K]EntryWithTimeout[V]),
	}
}

// Read returns the associated value of a key.
func (c *Cache[K, V]) Read(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var zeroV V

	e, ok := c.data[key]

	switch {
	case !ok:
		return zeroV, false
	case e.expires.Before(time.Now()):
		delete(c.data, key)
		return zeroV, false
	default:
		return e.value, true
	}
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = EntryWithTimeout[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}

	return nil
}

// Delete removes the entry of a given key.
func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}
