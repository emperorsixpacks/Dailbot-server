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
	items := make(map[string]Item)
	once.Do(func() {
		defaultCache = &cache{items: items}
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
	items   map[string]Item
	mu      sync.RWMutex
	janitor *Janitor
}

func (c *cache) Set(key string, value interface{}, duration time.Duration) (int, error) {
	var d int64
	if duration == 0 {
		duration = DefaultDuration
	}
	if duration > 0 {
		d = time.Now().Add(duration).UnixNano()
	}
	c.mu.RLock()
	c.items[key] = Item{
		value:    value,
		duration: d,
	}
	return 0, nil
}
func (c *cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, ok := c.items[key]
	c.mu.RUnlock()
	if !ok {
		return nil, false
	}
	return item.value, true
}

func (c *cache) Delete(key string) bool {
	c.mu.RLock()
	if _, ok := c.items[key]; !ok {
		return false
	}
	delete(c.items, key)
	c.mu.RUnlock()
	return true
}

func (c *cache) Len() int {
	c.mu.RLock()
	n := len(c.items)
	c.mu.RUnlock()
	return n
}

func (c *cache) Flush() {
	c.mu.RLock()
	c.items = map[string]Item{}
	c.mu.Unlock()
}

type Janitor struct {
	interval time.Duration
	stop     chan bool
}

func (j *Janitor) Run(c *cache) {
	ticker := time.NewTicker(j.interval)
	for {
		select {
		case <-ticker.C:
			// Do something
		case <-j.stop:
			ticker.Stop()
		}
	}
}

func stopExceution(c *cache) {
	c.janitor.stop <- true

}

// TODO we could design this better
func runJanitor(interval time.Duration, c *cache) {
	janitor := &Janitor{
		interval: interval,
		stop:     make(chan bool),
	}
	c.janitor = janitor
	janitor.Run(c)
}

// NOTE unix values and thier differrences
// TODO add increment for int
