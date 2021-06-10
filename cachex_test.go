package cachex

import (
	"reflect"
	"sync"
	"testing"
)

var (
	cacheInstance *Cache
	once          sync.Once
	C             *Cache
)

func init() {
	var c *Cache
	once.Do(func() {
		c = NewCache(1024)
		cacheInstance = new(Cache)
	})
	C = c
}

func TestCache_Set(t *testing.T) {
	values := [][]byte{[]byte("hello world01"), []byte("hello world02")}
	keys := [][]byte{[]byte("hello01"), []byte("hello02")}
	for k, v := range values {
		key := string(keys[k])
		C.Set(key, v)
		val, err := C.Get(key)
		if err != nil {
			t.Fatalf("Expect key:%v ,Value:%v ,Error:%v", key, v, err)
		}
		if !reflect.DeepEqual(val, v) {
			t.Fatalf("This key:%v,Expect value:%v, Get Value:%v", key, v, val)
		}
		t.Logf("value:%v ", val)

	}
}

func TestCache_Get(t *testing.T) {
	assertSuccessKey := "1234"
	assertSuccessValue := []byte{49, 50, 51, 52, 53}
	C.Set(assertSuccessKey, assertSuccessValue)
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  *Cache
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:   "Success",
			fields: C,
			args: args{
				key: assertSuccessKey,
			},
			want:    assertSuccessValue,
			wantErr: false,
		},
		{
			name:   "FailbyKeyExist",
			fields: C,
			args: args{
				key: "111111",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := C.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestCache_Get(t *testing.T) {
// 	c := getCache()
// 	assertSuccessKey := "1234"
// 	assertSuccessValue := []byte{49, 50, 51, 52, 53}
// 	c.Set(assertSuccessKey, assertSuccessValue)
// 	// type fields struct {
// 	// 	Mu        sync.RWMutex
// 	// 	CacheLen  int
// 	// 	Cache     map[interface{}]*list.Element
// 	// 	CacheList *list.List
// 	// }
// 	type args struct {
// 		key string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  Cache
// 		args    args
// 		want    []byte
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := c
// 			got, err := c.Get(tt.args.key)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Cache.Get() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Cache.Get() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
