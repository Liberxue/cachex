# cachex

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