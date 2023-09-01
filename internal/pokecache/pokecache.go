package pokecache

import (
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Entries map[string]cacheEntry
}

func NewCache() Cache {
	return Cache{
		Entries: make(map[string]cacheEntry),
	}
}

func (c Cache) Add(key string, val []byte) bool {
	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	if _, ok := c.Entries[key]; ok {
		return true
	}
	return false
}

func (c Cache) Get(key string) ([]byte, bool) {
	if entry, ok := c.Entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}
