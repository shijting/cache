package cache

import (
	"sync"
	"sync/atomic"
	"time"
)

type Cache struct {
	length int64
	maxMemory int64 // 最大内存 (字节)
	data map[string]interface{}
	rw sync.RWMutex
}
func NewCache() ICache {
	return &Cache{length: 0, data: make(map[string]interface{}), maxMemory: 10485760}
}

func (c *Cache) SetMaxMemory(size string) {
	panic("implement me")
}

func (c *Cache) Set(key string, val interface{}, exp time.Duration) {
	c.rw.Lock()
	defer func() {
		c.rw.Unlock()
	}()
	if !c.Exists(key) {
		atomic.AddInt64(&c.length, 1)
	}

	c.data[key] = val
	time.AfterFunc(exp, func() {
		delete(c.data, key)
		if c.length >0 {
			atomic.AddInt64(&c.length, -1)
		}
	})
}

func (c *Cache) Get(key string) (val interface{}, ok bool) {
	c.rw.RLock()
	defer c.rw.RUnlock()
	val, ok = c.data[key]
	return
}

func (c *Cache) Del(key string) bool {
	c.rw.Lock()
	defer c.rw.Unlock()
	delete(c.data, key)
	if c.length >0 {
		atomic.AddInt64(&c.length, -1)
	}
	return true
}

func (c *Cache) Exists(key string) bool {
	_,ok := c.data[key]
	return ok
}

func (c *Cache) Flush() bool {
	c.data = make(map[string]interface{})
	c.length = 0
	return true
}

func (c *Cache) Keys() int64 {
	return c.length
}


