package pokecache

import (
	"io"
	"net/http"
	"sync"
	"time"
)

type Cache struct {
	data     map[string]cacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		data:     make(map[string]cacheEntry),
		mu:       &sync.RWMutex{},
		interval: interval,
	}
	return &cache
}

// Create cacheEntry and insert it to map
func (c *Cache) Add(key string, val []byte) {
	cacheEntryTemp := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Lock()
	c.data[key] = cacheEntryTemp
	c.mu.Unlock()
}

// If this url data has already existed in cacheEntry
//
//	return cacheEntry data
//
// If this url data has not existed in cacheEntry yet
//
//	send request to this url
//	save response to cacheEntry
//	return data
func (c *Cache) Get(key string) ([]byte, bool, error) {
	c.mu.RLock()
	tempCacheEntry := c.data[key]
	c.mu.RLock()
	if tempCacheEntry.createdAt != (time.Time{}) {
		return tempCacheEntry.val, true, nil
	} else {
		req, err := http.NewRequest("GET", key, nil)
		if err != nil {
			return nil, false, err
		}
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return nil, false, err
		}
		defer res.Body.Close()
		byteRes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, false, err
		}
		c.mu.Lock()
		c.Add(key, byteRes)
		c.mu.Unlock()
		return byteRes, false, nil
	}
}

func (c *Cache) reapLoop()
