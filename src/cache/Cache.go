package cache

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Cache struct {
	// 缓存数量
	length int64 `ss:"-"`
	// 最大内存 (字节)
	maxMemory int `ss:"-"`
	// 线程安全的map
	//data sync.Map
	data map[string]interface{}
	lock sync.RWMutex
}



func NewCache() ICache {
	// 默认最大内存为10m
	return &Cache{length: 0, data: make(map[string]interface{}), maxMemory: defaultMemory}
}

func (c *Cache) SetMaxMemory(size string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	b := 0
	// 最大内存 转换成字节
	num, err := strconv.Atoi(size[:len(size)-2])
	fmt.Println("num", num)
	if err !=nil {
		return fmt.Errorf("parsing \"%s\": invalid syntax", size)
	}
	switch strings.ToLower(size[len(size) -2:]) {
	case "gb":
		b = num * 1024 * 1024 * 1024
	case "mb":
		b = num * 1024 * 1024
	case "kb":
		b = num * 1024
	default:
		b = defaultMemory
	}
	c.maxMemory = b
	return nil
}

func (c *Cache) Set(key string, val interface{}, exp time.Duration) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if SizeOf(c) > c.maxMemory {
		return fmt.Errorf("over the memory limit: %dByte \n", c.maxMemory)
	}
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
	return nil
}

func (c *Cache) Get(key string) (val interface{}, ok bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	val, ok = c.data[key]
	return
}

func (c *Cache) Del(key string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
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


