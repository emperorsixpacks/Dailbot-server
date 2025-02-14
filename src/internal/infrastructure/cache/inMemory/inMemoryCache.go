package inmemoryCache

import (
	"sync"
	"time"
)

var (
	once         sync.Once
	defaultCache *cache
)

func New() *cache {
	once.Do(func() {
		defaultCache = &cache{}
	})
	return defaultCache
}

type Item struct {
	value      interface{}
	expiration string
}

type cache struct {
	items map[string]Item
	mu    sync.RWMutex
}

func (c *cache) Set(key string, value interface{}, duration time.Duration) (int, error) {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.items[key] = Item{
		value:      value,
		expiration: "hello",
	}
	return 0, nil
}
func (c *cache) Get(key string) (interface{}, bool) {
	defer c.mu.Unlock()
	c.mu.Lock()
	item, ok := c.items[key]
	if !ok {
		return nil, false
	}
	return item.value, true
}
