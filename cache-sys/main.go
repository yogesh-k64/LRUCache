package main

import (
	"container/list"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/rs/cors"
)

type CacheItem struct {
	key        string
	value      interface{}
	expiration int64 // Unix timestamp
}

type RecentCache struct {
	capacity int
	cache    map[string]*list.Element
	eviction *list.List
	lock     sync.Mutex
}

func NewLRUCache(capacity int) *RecentCache {
	cache := &RecentCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		eviction: list.New(),
	}

	go cache.startEviction()

	return cache
}

func (c *RecentCache) startEviction() {
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		c.removeExpiredItems()
	}
}

func (c *RecentCache) removeExpiredItems() {
	c.lock.Lock()
	defer c.lock.Unlock()

	now := time.Now().Unix()
	for {
		element := c.eviction.Back()
		if element == nil {
			break
		}
		item := element.Value.(*CacheItem)
		if item.expiration > 0 && now > item.expiration {
			c.eviction.Remove(element)
			delete(c.cache, item.key)
		} else {
			break
		}
	}
}

func (c *RecentCache) Get(key string) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if element, found := c.cache[key]; found {
		item := element.Value.(*CacheItem)
		if time.Now().Unix() > item.expiration {
			c.eviction.Remove(element)
			delete(c.cache, key)
			return nil, false
		}
		c.eviction.MoveToFront(element)
		return item.value, true
	}
	return nil, false
}

func (c *RecentCache) Set(key string, value interface{}, expiration time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if element, found := c.cache[key]; found {
		c.eviction.MoveToFront(element)
		element.Value.(*CacheItem).value = value
		element.Value.(*CacheItem).expiration = time.Now().Add(expiration).Unix()
	} else {
		if c.eviction.Len() >= c.capacity {
			c.evict()
		}
		item := &CacheItem{
			key:        key,
			value:      value,
			expiration: time.Now().Add(expiration).Unix(),
		}
		element := c.eviction.PushFront(item)
		c.cache[key] = element
	}
}

func (c *RecentCache) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if element, found := c.cache[key]; found {
		c.eviction.Remove(element)
		delete(c.cache, key)
	}
}

func (c *RecentCache) evict() {
	element := c.eviction.Back()
	if element != nil {
		c.eviction.Remove(element)
		item := element.Value.(*CacheItem)
		delete(c.cache, item.key)
	}
}

var cache *RecentCache

func main() {
	cache = NewLRUCache(5)

	mux := http.NewServeMux()
	mux.HandleFunc("/get", getHandler)
	mux.HandleFunc("/set", setHandler)
	mux.HandleFunc("/delete", deleteHandler)

	handler := cors.Default().Handler(mux)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(PORT, handler)
}
