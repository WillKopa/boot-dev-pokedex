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
	interval		time.Duration
	cach_map		map[string]cacheEntry
	sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	new_cache := Cache {
		cach_map: map[string]cacheEntry{},
	}
	go new_cache.reapLoop(interval)
	return &new_cache
}

func (c *Cache) Add (key string, val []byte) {
	c.Lock()
	c.cach_map[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.Unlock()
}

func (c *Cache) Get (key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()
	cm, ok := c.cach_map[key]
	if !ok {
		return nil, ok
	}
	return cm.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Lock()
		for key, data := range (c.cach_map) {
			if data.createdAt.Add(interval).Before(time.Now()) {
				delete(c.cach_map, key)
			}
		}
		c.Unlock()
	}
}