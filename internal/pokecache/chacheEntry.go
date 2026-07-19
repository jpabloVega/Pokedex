package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	ent := Cache{
		entry: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go ent.reapLoop(interval)

	return ent
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entry[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.entry[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	tick := time.NewTicker(interval)
	for range tick.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.entry {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entry, k)
		}
	}
}
