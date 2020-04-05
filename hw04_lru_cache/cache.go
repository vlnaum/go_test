package hw04_lru_cache //nolint:golint,stylecheck

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*listItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if item, ok := lc.items[key]; ok {
		item.Value = cacheItem{key, value}
		lc.queue.MoveToFront(item)
		return true
	}

	if lc.capacity == lc.queue.Len() {
		delete(lc.items, lc.queue.Back().Value.(cacheItem).key)
		lc.queue.Remove(lc.queue.Back())
	}

	element := lc.queue.PushFront(cacheItem{key, value})
	lc.items[key] = element
	return false
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if item, ok := lc.items[key]; ok {
		lc.queue.MoveToFront(item)
		return item.Value.(cacheItem).value, true
	}

	return nil, false
}

func (lc *lruCache) Clear() {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.queue = NewList()
	lc.items = make(map[Key]*listItem, lc.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*listItem, capacity)}
}
