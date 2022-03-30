package fifo

import (
	"sync"

	list "github.com/bahlo/generic-list-go"
)

type FIFO[K comparable, V any] struct {
	m     sync.Mutex
	ll    *list.List[*entry[K, V]]
	cache map[K]*list.Element[*entry[K, V]]
	size  int
}

type entry[K comparable, V any] struct {
	key   K
	value *V
}

type Evicted[K comparable, V any] struct {
	Key   K
	Value V
}

func (L *FIFO[K, V]) Get(key K) *V {
	L.m.Lock()
	defer L.m.Unlock()

	if e, ok := L.cache[key]; ok {
		return e.Value.value
	}

	return nil
}

func (L *FIFO[K, V]) Push(key K, value V) *Evicted[K, V] {
	if L.size < 1 {
		return &Evicted[K, V]{key, value}
	}

	L.m.Lock()
	defer L.m.Unlock()

	if e, ok := L.cache[key]; ok {
		e.Value.value = &value
		return nil
	}

	e := L.ll.Back()
	i := e.Value
	evictedKey := i.key
	evictedValue := i.value
	delete(L.cache, i.key)

	i.key = key
	i.value = &value
	L.cache[key] = e
	L.ll.MoveToFront(e)
	if evictedValue != nil {
		return &Evicted[K, V]{evictedKey, *evictedValue}
	}
	return nil
}

func (L *FIFO[K, V]) Len() int {
	L.m.Lock()
	defer L.m.Unlock()

	return len(L.cache)
}

func (L *FIFO[K, V]) Remove(key K) *V {
	L.m.Lock()
	defer L.m.Unlock()

	if e, ok := L.cache[key]; ok {
		value := e.Value.value
		L.ll.MoveToBack(e)
		e.Value.value = nil
		delete(L.cache, key)
		return value
	}

	return nil
}

func (L *FIFO[K, V]) Victim() *K {
	if L.size < 1 {
		return nil
	}

	L.m.Lock()
	defer L.m.Unlock()

	e := L.ll.Back()
	i := e.Value
	evictedKey := i.key
	evictedValue := i.value

	if evictedValue == nil {
		return nil
	}

	return &evictedKey
}

func New[K comparable, V any](size int) *FIFO[K, V] {
	c := &FIFO[K, V]{
		ll:    list.New[*entry[K, V]](),
		cache: make(map[K]*list.Element[*entry[K, V]], size),
		size:  size,
	}

	for i := 0; i < size; i++ {
		c.ll.PushBack(&entry[K, V]{})
	}

	return c
}
