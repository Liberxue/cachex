package cachex

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"testing"
)

var (
	c    *Cache
	once sync.Once
)

func init() {
	once.Do(func() {
		c = NewCache(102400, 2)
		_ = c.startCleanExpireOldestCache
	})
}

func TestCache_Set(t *testing.T) {
	values := [][]byte{[]byte("hello world01"), []byte("hello world02")}
	keys := [][]byte{[]byte("hello01"), []byte("hello02")}
	for k, v := range values {
		key := string(keys[k])
		c.Set(key, v)
		val, err := c.Get(key)
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
	c.Set(assertSuccessKey, assertSuccessValue)
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
			fields: c,
			args: args{
				key: assertSuccessKey,
			},
			want:    assertSuccessValue,
			wantErr: false,
		},
		{
			name:   "FailbyKeyExist",
			fields: c,
			args: args{
				key: "111111",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Get(tt.args.key)
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
		c.Set(key, v)
		val, err := c.Get(key)
		c.cleanExpireOldestCache()
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

func BenchmarkCache_Set(b *testing.B) {
	cacheInstance := NewCache(10e6, 1)
	for n := 0; n < b.N; n++ {
		cacheInstance.Set(fmt.Sprint(n%1000000), []byte("value"))
		cacheInstance.Get(fmt.Sprint(n % 1000000))
	}
}
