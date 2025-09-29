package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt     time.Time
	CacheEntryVal []byte
}

type Cache struct {
	CacheEntries  map[string]cacheEntry
	CacheInterval time.Duration
	mu            sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	defaultMap := make(map[string]cacheEntry)
	// mutex doesn't need instiated
	cache := Cache{CacheEntries: defaultMap, CacheInterval: interval}
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	//entries, ok := c.CacheEntries

	c.CacheEntries[key] = cacheEntry{CreatedAt: time.Now(), CacheEntryVal: val}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.CacheEntries[key]
	if !ok {
		return []byte{}, false
	}
	return entry.CacheEntryVal, true
}

func (c *Cache) ReadLoop(interval time.Duration) {
	c.mu.Lock()
	for key, value := range c.CacheEntries {
		c.mu.Lock()
		created := value.CreatedAt
		exitTime := created.Add(interval)
		now := time.Now()
		comparison := now.Compare(exitTime)
		if comparison == 0 || comparison == 1 {
			delete(c.CacheEntries, key)
		}
	}
	c.mu.Unlock()
}
