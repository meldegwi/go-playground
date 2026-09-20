package cache

// Cache is a gerneric cache, its key can be any comparale value and the attached,
// value can be anything.
type Cache[K comparable, V any] struct {
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
	v, found := c.data[key]
	return v, found
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.data[key] = value

	return nil
}
