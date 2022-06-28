package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	lru "github.com/hashicorp/golang-lru"
)

var cache Cache

type Cache struct {
	sync.Mutex
	*lru.Cache
	once sync.Once
	size int
}

func InitCache(size int) error {
	var err error
	cache.once.Do(func() {
		cache.Cache, err = lru.New(size)
		if err != nil {
			return
		}
		cache.size = size
	})
	return err
}

func (c *Cache) Resize(size int) {
	c.Lock()
	defer c.Unlock()

	c.Cache.Resize(size)
	c.size = size
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" || (r.Method != http.MethodPost && r.Method != http.MethodPut) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cache.Add(key, body)
	w.WriteHeader(http.StatusAccepted)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, ok := cache.Get(key)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bin, _ := json.Marshal(value)
	w.Write(bin)
}

func ResizeHandler(w http.ResponseWriter, r *http.Request) {
	size := r.URL.Query().Get("size")
	if size == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resize, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cache.Resize(int(resize))
}

func ClearHandler(w http.ResponseWriter, r *http.Request) {
	cache.Purge()
}
