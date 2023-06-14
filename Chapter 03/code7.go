package main

import "fmt"

// contains code for handling a hashed map data structure (self built - not inbuilt golang's version)
// naive approach for hashed map data structure

type HashedMaps struct {
	items [100]string
}

func NewHashedMaps() HashedMaps {
	return HashedMaps{items: [100]string{}}
}

func (h HashedMaps) GetHash(key string) int {
	totalSum := 0
	for _, v := range key {
		totalSum = totalSum + int(v)
	}
	hashKey := totalSum % 100
	return hashKey
}

func (h *HashedMaps) Set(key, val string) {
	hashedKey := h.GetHash(key)
	h.items[hashedKey] = val
}

func (h *HashedMaps) Get(key string) string {
	hashedKey := h.GetHash(key)
	return h.items[hashedKey]
}

func main() {
	aa := NewHashedMaps()
	aa.Set("aa", "sample value")
	fmt.Println(aa.Get("aa"))

	fmt.Println(aa.GetHash("a"))
	fmt.Println(aa.GetHash("ABB"))

	aa.Set("a", "sample value")
	aa.Set("ABB", "unexpected sample value")
	fmt.Println(aa.GetHash("a"))

}
