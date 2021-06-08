package cachex

import (
	"container/list"
	"errors"
	"sync"
	"time"
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
	body   []byte //snappy / gzip Compress
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
	// check Cache len
	if c.CacheSize <= len(key)+len(value) || c.CacheSize == 0 {
		return errors.New("Cache is full")
	}
	// Expire oldest form Cache
	if c.CacheList.Len() > c.CacheSize {
		// clean ....
	}
	// assert key is existed...
	if element, existed := c.Cache[key]; existed {
		//c.CacheList.PushBack(element)
		// the left side data is best new data ... lru
		c.CacheList.MoveToFront(element)
		element.Value.(*baseCache).body = value
		return nil
	}
	// key not exist
	element := c.CacheList.PushFront(
		&baseCache{
			header: &header{
				ttl:        0,
				key:        key,
				tag:        "",
				createTime: time.Now().UnixNano(),
			},
			body: value,
		},
	)
	c.Cache[key] = element
	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	if element, ok := c.Cache[key]; ok {
		c.CacheList.MoveToFront(element) //lru
		return element.Value.(*baseCache).body, nil
	}
	return nil, errors.New("The key is exist")
}
