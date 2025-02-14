package inmemoryCache

import (
	"sync"
	"time"
)

const (
	DefaultDuration time.Duration = 0
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
	value    interface{}
	duration int64
}

func (this Item) Expired() bool {
	if this.duration == 0 {
		return false
	}
	return time.Now().UnixNano() > this.duration
}

type cache struct {
	items map[string]Item
	mu    sync.RWMutex
}

func (c *cache) Set(key string, value interface{}, duration time.Duration) (int, error) {
	defer c.mu.Unlock()
	var d int64
	if duration == 0 {
		duration = DefaultDuration
	}
	if duration > 0 {
		d = time.Now().Add(duration).UnixNano()
	}
	c.mu.Lock()
	c.items[key] = Item{
		value:    value,
		duration: d,
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

// NOTE unix values and thier differrences
