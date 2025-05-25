package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) (Cache, error) {
	cache := Cache{
		mu:    sync.Mutex{},
		cache: map[string]cacheEntry{},
	}
	go cache.reapLoop(interval)
	return cache, nil
}

func (c *Cache) Add(key string, value []byte) error {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = entry
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
