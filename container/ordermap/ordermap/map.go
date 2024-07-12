package ordermap

import "github.com/smallnest/exp/container/list"

type OrderedMap[K comparable, V any] struct {
	entries map[K]*Entry[K, V]

	list *list.List[*Entry[K, V]]
}

func (m *OrderedMap[K, V]) Set(key K, value V) (val V, existed bool) {
	if entry, existed := m.entries[key]; existed { // 如果key存在，就是更新
		oldValue := entry.Value
		entry.Value = value
		return oldValue, true
	}
	entry := &Entry[K, V]{
		Key:   key,
		Value: value,
	}
	entry.element = m.list.PushBack(entry) // 加入到链表
	m.entries[key] = entry                 // 加入到map中
	return value, false
}
func (m *OrderedMap[K, V]) Delete(key K) (val V, exist bool) {
	entry, exist := m.entries[key]
	if exist { // 如果存在
		m.list.Remove(entry.element) // 从链表中移除
		delete(m.entries, key)       // 从map中删除
		return entry.Value, true
	}
	return
}
func (m *OrderedMap[K, V]) Get(key K) (val V, exist bool) {
	entry, exist := m.entries[key]
	if exist {
		return entry.Value, true
	}
	return
}
func (m *OrderedMap[K, V]) Range(f func(key K, value V) bool) {
	for e := m.list.Front(); e != nil; e = e.Next() {
		if e.Value != nil {
			if ok := f(e.Value.Key, e.Value.Value); !ok {
				return
			}
		}
	}
}

type Entry[K comparable, V any] struct {
	Key     K
	Value   V
	element *list.Element[*Entry[K, V]]
}

func (e *Entry[K, V]) Next() *Entry[K, V] {
	entry := e.element.Next()
	if entry == nil {
		return nil
	}
	return entry.Value
}
func (e *Entry[K, V]) Prev() *Entry[K, V] {
	entry := e.element.Prev()
	if entry == nil {
		return nil
	}
	return entry.Value
}
