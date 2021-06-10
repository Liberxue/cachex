package cachex

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

// ExpireCallBackNotice -> Expire CallBack Notice
type ExpireCallBackNotice func(key string, value []byte) bool

// CheckExpireCallback -> Check Expire CallBack Notice
type CheckExpireCallback func(key string, value []byte)

// Cachex interface
type Cachex interface {
	Set(key string, value []byte)
	Get(key string) ([]byte, error)
}

// Cache struct
type Cache struct {
	Mu        sync.RWMutex
	TTL       uint
	CacheLen  uint
	Cache     map[interface{}]*list.Element
	CacheList *list.List
}

// Cache Protocol
type baseCache struct {
	header *header
	body   []byte //snappy / gzip Compress
}

// Cache Header
type header struct {
	key, tag   string
	TTL        uint
	createTime int64
}

//NewCache Cache
//input cacheLen uint
//return *Cache
func NewCache(cacheLen, ttl uint) *Cache {
	return &Cache{
		CacheLen:  cacheLen,
		TTL:       ttl,
		CacheList: list.New().Init(), //clean list && init list
		Cache:     make(map[interface{}]*list.Element, cacheLen),
	}
}

//Set Cache
//input key string,value []byte
//return error
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
				TTL:        0,
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

//Get Cache
//input key string
//return []byte,error
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

// cleanExpireOldestCacheByTTL
func (c *Cache) cleanExpireOldestCacheByTTL(nowTime int64) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	// o(n)
	for element := c.CacheList.Back(); element != nil; element = element.Next() {
		val := element.Value.(*baseCache)
		if int64(val.header.TTL)+val.header.createTime < nowTime {
			delete(c.Cache, val.header.key)
		}
	}
}

// start Clean ExpireOldest Cache Job...
func (c *Cache) startCleanExpireOldestCache() {
	if c.TTL <= 1 {
		return
	}
	timer := time.NewTimer(time.Second)
	go func(timer *time.Timer) {
		c.cleanExpireOldestCacheByTTL(time.Now().UnixNano())
		<-timer.C
	}(timer)
	timer.Reset(0 * time.Second)
	time.Sleep(time.Second * time.Duration(c.TTL))
}
