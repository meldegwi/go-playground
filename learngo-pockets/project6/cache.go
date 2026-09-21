package cache

import (
	"slices"
	"sync"
	"time"
)

// Cache is a generic cache, its key can be any comparale value and the attached,
// value can be anything.
type Cache[K comparable, V any] struct {
	ttl time.Duration

	mu   sync.Mutex
	data map[K]EntryWithTimeout[V]

	mxSize            int
	chronologicalKeys []K
}

type EntryWithTimeout[V any] struct {
	value   V
	expires time.Time // after that time the value is expired.
}

// New makes a new instance of the Cache struct.
func New[K comparable, V any](cacheSize int, ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:               ttl,
		data:              make(map[K]EntryWithTimeout[V]),
		mxSize:            cacheSize,
		chronologicalKeys: make([]K, 0, cacheSize),
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
		c.deleteKeyValue(key)
		return zeroV, false
	default:
		return e.value, true
	}
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, valueFound := c.data[key]

	switch {
	case valueFound:
		c.deleteKeyValue(key)
	case len(c.data) == c.mxSize:
		c.deleteKeyValue(c.chronologicalKeys[0])
	}

	c.addKeyValue(key, value)
}

// Delete removes the entry of a given key.
func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.deleteKeyValue(key)
}

// addKeyValue inserts a key and its value into the cache.
func (c *Cache[K, V]) addKeyValue(key K, value V) {
	c.data[key] = EntryWithTimeout[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}

	c.chronologicalKeys = append(c.chronologicalKeys, key)
}

// deleteKeyValue removes a key and its value from the cache.
func (c *Cache[K, V]) deleteKeyValue(key K) {
	c.chronologicalKeys = slices.DeleteFunc(
		c.chronologicalKeys,
		func(k K) bool {
			return k == key
		})

	delete(c.data, key)
}
