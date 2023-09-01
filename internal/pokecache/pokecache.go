package pokecache

import (
	"fmt"
	"time"
)

// I used a Cache struct to hold a map[string][cacheEntry] and a mutex to protect the map across goroutines. A cacheEntry should be a struct with two fields:

// createdAt - A time.Time that represents when the entry was created.
// val - A []byte that represents the raw data we're caching.
// You'll probably want to expose a NewCache() function that creates a new cache with a configurable interval (time.Duration).

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

func ThisIsATest() {
	fmt.Println("testing that i can import an internal package")
}
