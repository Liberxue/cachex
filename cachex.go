package cachex

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

type ExpireCallBackNotice func(key string, value []byte) bool

type CheckExpireCallback func(key string, value []byte)
type Cachex interface {
	Set(key string, value []byte)
	Get(key string) ([]byte, error)
}

type Cache struct {
	Mu        sync.RWMutex
	TTl       uint
	CacheLen  uint
	Cache     map[interface{}]*list.Element
	CacheList *list.List
}

type baseCache struct {
	header *header
	body   []byte //snappy / gzip Compress
}
type header struct {
	key, tag   string
	ttl        uint
	createTime int64
}

// New Cache init list
// @input cacheLen int
// @return *Cache
func NewCache(cacheLen uint) *Cache {
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
	if c.CacheLen <= 0 {
		return errors.New("cacheLen is 0")
	}
	go c.startCleanExpireOldestCache()
	// assert key is existed...
	if element, existed := c.Cache[key]; existed {
		//c.CacheList.PushBack(element)
		// the left side data is best new data ... lru
		c.CacheList.MoveToFront(element)
		element.Value.(*baseCache).body = value
		return nil
	}
	// Expire oldest form Cache
	if uint(c.CacheList.Len()) > c.CacheLen {
		// clean ....
		c.cleanExpireOldestCache()
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

//
func (c *Cache) cleanExpireOldestCacheByTTl(nowTime int64) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	// o(n)
	for element := c.CacheList.Back(); element != nil; element = element.Next() {
		val := element.Value.(*baseCache)
		if int64(val.header.ttl)+val.header.createTime < nowTime {
			delete(c.Cache, val.header.key)
		}
	}
}

// start Clean ExpireOldest Cache Job...
func (c *Cache) startCleanExpireOldestCache() {
	timer := time.NewTimer(time.Second)
	go func() {
		c.cleanExpireOldestCacheByTTl(time.Now().UnixNano())
		<-timer.C
	}()
	timer.Reset(0 * time.Second)
	time.Sleep(time.Second * time.Duration(c.TTl))
}
