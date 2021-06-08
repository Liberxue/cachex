package cachex

import (
	"container/list"
	"sync"
)

type Cachex interface {
	Set(key string, value []byte)
	Get(key string) ([]byte, error)
}

type Cache struct {
	Mu        sync.RWMutex
	CacheSize int
	Cache     map[interface{}]*list.Element
	CacheList *list.List
}

type baseCache struct {
	header *header
	body   []byte
}
type header struct {
	key, tag   string
	ttl        int
	createTime int64
}

func NewCache(cacheSize int) *Cache {
	return &Cache{
		CacheSize: cacheSize,
		CacheList: list.New().Init(), //clean list && init list
		Cache:     make(map[interface{}]*list.Element, cacheSize),
	}
}

func (c *Cache) Set(key string, value []byte) error {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	if c.Cache == nil {
		c.Cache = make(map[interface{}]*list.Element, c.CacheSize)
		c.CacheList = list.New() //init list
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	return nil, nil
}
