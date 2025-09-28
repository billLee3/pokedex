package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct{
	createdAt time.Time
	cacheEntryVal []byte
}

type Cache struct{
	cacheEntries map[string]cacheEntry
	cacheInterval time.Duration
	mu sync.Mutex
}

func NewCache(interval time.Duration) Cache{
	defaultMap := make(map[string]cacheEntry{})
	// mutex doesn't need instiated
	cache := Cache{cacheEntries: defaultMap, cacheInterval: interval}
	return cache
}

func (c Cache) Add(key string, val []byte) error{
	c.mu.Lock()
	if err := c.cacheEntries; err != nil{
		return err
	}
	c.cacheEntries[key] = cacheEntry{createdAt: time.Now(), cacheEntryVal: val}
	c.mu.Unlock()

	return nil
}

func (c Cache) Get(key string) ([]byte, bool){
	entry, ok := c.cacheEntries[key]
	if !ok{
		return []byte{}, false
	}
	return entry.cacheEntryVal, true
}
