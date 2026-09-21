package cache

import "sync"

// Cache is a gerneric cache, its key can be any comparale value and the attached,
// value can be anything.
type Cache[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// New makes a new instance of the Cache struct.
func New[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
		data: make(map[K]V),
	}
}

// Read returns the associated value of a key.
func (c *Cache[K, V]) Read(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, found := c.data[key]

	return v, found
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value

	return nil
}

// Delete removes the entry of a given key.
func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}
