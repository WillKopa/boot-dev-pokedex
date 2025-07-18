package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt		time.Time
	val				[]byte
}

type Cache struct {
	cache		map[string]cacheEntry
	sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	new_cache := Cache {
		cache: map[string]cacheEntry{},
	}
	go new_cache.reapLoop(interval)
	return &new_cache
}

func (c *Cache) Add (key string, val []byte) {
	c.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.Unlock()
}

func (c *Cache) Get (key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()
	cm, ok := c.cache[key]
	if !ok {
		return nil, ok
	}
	return cm.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Lock()
		for key, data := range (c.cache) {
			if data.createdAt.Add(interval).Before(time.Now()) {
				delete(c.cache, key)
			}
		}
		c.Unlock()
	}
}