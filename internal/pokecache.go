package internal

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	var mu1 sync.Mutex
	cache := Cache{
		cache:    map[string]cacheEntry{},
		mu:       mu1,
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}

func (c Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cache[key] = entry
}

func (c Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	if ok {
		return entry.val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
