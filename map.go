package syncmap

import "sync"

// Map is a generics wrapper over the sync.Map from Go standard library.
type Map[K comparable, V any] struct {
	v sync.Map
}

// Clear deletes all the entries, resulting in an empty Map.
func (m *Map[K, V]) Clear() {
	m.v.Clear()
}

// Contains returns true if key exists in the map.
func (m *Map[K, V]) Contains(key K) bool {
	_, ok := m.v.Load(key)
	return ok
}

// Delete deletes the value for a key.
func (m *Map[K, V]) Delete(key K) {
	m.v.Delete(key)
}

// Load returns the value stored in the map for a key, or zero value if key is
// not present. The ok result indicates whether value was found in the map.
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.v.Load(key)
	if !ok {
		return value, ok
	}
	return v.(V), ok
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(key K, value V) {
	m.v.Store(key, value)
}

// LoadAndDelete deletes the value for a key, returning the previous value if
// any. The loaded result reports whether the key was present.
func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.v.LoadAndDelete(key)
	if !loaded {
		return value, loaded
	}
	return v.(V), loaded
}

// LoadOrStore returns the existing value for the key if present. Otherwise, it
// stores and returns the given value. The loaded result is true if the value
// was loaded, false if stored.
func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, loaded := m.v.LoadOrStore(key, value)
	return a.(V), loaded
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
func (m *Map[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.v.CompareAndDelete(key, old)
}

// CompareAndSwap swaps the old and new values for key if the value stored in
// the map is equal to old.
func (m *Map[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.v.CompareAndSwap(key, old, new)
}

// Range calls f sequentially for each key and value present in the map. If f
// returns false, range stops the iteration. See Map.Range documentation sync
// package for more details.
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.v.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}
