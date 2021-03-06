# cachex

[![GitHub Action Status](https://github.com/Liberxue/cachex/workflows/Tests/badge.svg)](https://github.com/Liberxue/cachex/actions?query=workflow%3ATests)
[![Build status](https://circleci.com/gh/Liberxue/cachex/tree/main.svg?style=shield&circle-token=14b3bd789e93258fc2e7a90a26337db82ce20514)](https://circleci.com/gh/Liberxue/cachex)
[![Go Report Card](https://goreportcard.com/badge/github.com/Liberxue/cachex)](https://goreportcard.com/report/github.com/Liberxue/cachex)
[![GoDoc](https://godoc.org/github.com/Liberxue/cachex?status.svg)](https://godoc.org/github.com/Liberxue/cachex)
[![Coverage Status](https://coveralls.io/repos/github/Liberxue/cachex/badge.svg)](https://coveralls.io/github/Liberxue/cachex)
[![license](https://img.shields.io/github/license/Liberxue/cachex.svg?maxAge=2592000)](https://github.com/Liberxue/cachex/blob/main/LICENSE)

## snappy example 
```go
var (
	c    *Cache
	once sync.Once
)

func init() {
	once.Do(func() {
		c = NewCache(102400)
	})
}
func main() {
	src := []byte(`{"test":"test"}`)
	encoded := snappy.Encode(nil, src)
	key := ("hello")
	err := c.Set(key, encoded)
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.Get(key)
	if err != nil {
		log.Fatal(err)
	}
	decoded, err := snappy.Decode(nil, encoded)
	if err != nil {
		log.Fatal(err)
	}
	valueStr := unsafe.Pointer(&decoded)
	fmt.Println(*(*string)(valueStr), key)
}
```