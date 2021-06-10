# cachex

[![GitHub Action Status](https://github.com/Liberxue/cachex/workflows/Tests/badge.svg)](https://github.com/Liberxue/cachex/actions?query=workflow%3ATests)
[![Build status](https://circleci.com/gh/Liberxue/cachex/tree/main.svg?style=shield&circle-token=14b3bd789e93258fc2e7a90a26337db82ce20514)](https://circleci.com/gh/Liberxue/cachex)
[![GoDoc](https://godoc.org/github.com/Liberxue/cachex?status.svg)](https://godoc.org/github.com/Liberxue/cachex)

## snappy example 
```go
func main() {
	src := []byte(`{"test":"test"}`)
	encoded := snappy.Encode(nil, src)
	c := cachex.NewCache(102400)
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