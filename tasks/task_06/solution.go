package main

import "container/list"

type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element, capacity),
	}
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	if c.capacity <= 0 {
		return
	}

	mapV, mapOk := c.items[key]
	if !mapOk {
		return
	}

	c.ll.MoveToFront(mapV)
	value = mapV.Value.(*entry[K, V]).value
	ok = true
	return
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity <= 0 {
		return
	}

	_, ok := c.items[key]
	if ok {
		c.items[key].Value.(*entry[K, V]).value = value
		return
	}

	if c.ll.Len() == c.capacity {
		oldValue := c.ll.Remove(c.ll.Back()).(*entry[K, V])
		delete(c.items, oldValue.key)
	}

	newElement := c.ll.PushFront(&entry[K, V]{
		key:   key,
		value: value,
	})

	c.items[key] = newElement
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
