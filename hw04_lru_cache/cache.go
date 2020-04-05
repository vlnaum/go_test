package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Capacity int
	Queue    List
	Items    map[Key]*listItem
}

type cacheItem struct {
	Key
	Value interface{}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := lc.Items[key]; ok {
		item.Value = cacheItem{key, value}
		lc.Queue.MoveToFront(item)
		return true
	}
	if lc.Capacity == lc.Queue.Len() {
		delete(lc.Items, lc.Queue.Back().Value.(cacheItem).Key)
		lc.Queue.Remove(lc.Queue.Back())
	}
	e := lc.Queue.PushFront(cacheItem{key, value})
	lc.Items[key] = e
	return false
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := lc.Items[key]; ok {
		lc.Queue.MoveToFront(item)
		return item.Value.(cacheItem).Value, true
	}
	return nil, false
}

func (lc *lruCache) Clear() {
	lc.Queue = NewList()
	lc.Items = make(map[Key]*listItem, lc.Capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{capacity, NewList(), make(map[Key]*listItem, capacity)}
}
