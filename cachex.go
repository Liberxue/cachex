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
	CacheLen  int
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

// New Cache init list
//
// @return *Cache
func NewCache(cacheLen int) *Cache {
	// go startCleanExpireOldestCache()
	return &Cache{
		CacheLen:  cacheLen,
		CacheList: list.New().Init(), //clean list && init list
		Cache:     make(map[interface{}]*list.Element, cacheLen),
	}
}

// SetCache
// @input key string,value []byte
// return error
func (c *Cache) Set(key string, value []byte) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if c.Cache == nil {
		c.Cache = make(map[interface{}]*list.Element, c.CacheLen)
		c.CacheList = list.New() //init list
	}
	// check Cache len
	if c.CacheLen == 0 {
		return errors.New("cacheLen is 0")
	}
	// Expire oldest form Cache
	if c.CacheList.Len() > c.CacheLen {
		// clean ....
		c.cleanExpireOldestCache()
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

// GetCacheByKey
// @input key string
// return []byte,error
func (c *Cache) Get(key string) ([]byte, error) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	if element, ok := c.Cache[key]; ok {
		c.CacheList.MoveToFront(element) //lru
		return element.Value.(*baseCache).body, nil
	}
	return nil, errors.New("key is exist")
}

// cleanExpireOldestCache
func (c *Cache) cleanExpireOldestCache() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	element := c.CacheList.Back()
	if element != nil {
		c.CacheList.Remove(element)
		key := element.Value.(*baseCache).header.key
		delete(c.Cache, key)
	}
}
