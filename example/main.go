package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Liberxue/cachex"
	"github.com/golang/snappy"
)

var (
	c    *cachex.Cache
	once sync.Once
)

func init() {
	once.Do(func() {
		c = cachex.NewCache(102400)
	})
}
func main() {
	src := []byte(`{"test":"test"}`)
	encoded := snappy.Encode(nil, src)
	var ch chan int
	ticker := time.NewTicker(time.Microsecond * 500)
	go func() {
		for range ticker.C {
			key := fmt.Sprintf("hello_%d", time.Now().UnixNano())
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
			fmt.Println(decoded)
		}
		ch <- 1
	}()
	<-ch
}
