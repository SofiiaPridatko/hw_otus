package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := c.items[key]; ok {
		entry := item.Value.(cacheItem)
		entry.value = value
		item.Value = entry
		c.queue.MoveToFront(item)

		return true
	}

	item := c.queue.PushFront(cacheItem{key: key, value: value})
	c.items[key] = item

	if c.queue.Len() > c.capacity {
		last := c.queue.Back()
		entry := last.Value.(cacheItem)
		delete(c.items, entry.key)
		c.queue.Remove(last)
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	c.queue.MoveToFront(item)
	entry := item.Value.(cacheItem)

	return entry.value, true
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
