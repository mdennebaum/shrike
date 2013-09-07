package cache

import (
	"time"
)

type Cache struct {
	sync.RWMutex
	Data           map[string]*Item
	DefaultTimeout time.Duration
}

func NewCache(defaultTimeout time.Duration) *Cache {
	if defaultTimeout == nil {
		defaultTimeout = (24 * time.Hour) * 30
	}
	c := Cache{make(map[string]*Item), defaultTimeout}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	val, ok := c.Data[key]
	c.RUnlock()
}

func (c *Cache) Set(key string, value interface{}, timeout time.Duration) {
	if timeout == nil {
		timeout = c.DefaultTimeout
	}
	c.Lock()
	c.Data[key] = &Item{value, timeout}
	c.Unlock()
}

func (c *Cache) Delete(key string) {
	c.Lock()
	delete(c.Data, key)
	c.Unlock()
}
