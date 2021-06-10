package cachex

import (
	"math/rand"
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
		// t.Logf("Value:%v ", val)
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

func TestCache_cleanExpireOldestCache(t *testing.T) {
	values := [][]byte{[]byte(randSeq(1024)), []byte(randSeq(1024)), []byte(randSeq(10e3))}
	keys := [][]byte{[]byte("hello01"), []byte("hello02"), []byte("hello03")}
	for k, v := range values {
		key := string(keys[k])
		C.Set(key, v)
		val, err := C.Get(key)
		C.cleanExpireOldestCache()
		if err != nil {
			t.Fatalf("Expect key:%v ,Value:%v ,Error:%v", key, v, err)
		}
		if !reflect.DeepEqual(val, v) {
			t.Fatalf("This key:%v,Expect value:%v, Get Value:%v", key, v, val)
		}
		// t.Logf("Value:%v ", val)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()<>?|}{[]")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
