package utils

import "sync"

// 非并发安全的 Set
type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(key ...T) {
	for _, key := range key {
		s.data[key] = struct{}{}
	}
}

func (s *Set[T]) Del(key ...T) {
	for _, key := range key {
		delete(s.data, key)
	}
}

func (s *Set[T]) Exist(key T) bool {
	_, ok := s.data[key]
	return ok
}

func (s *Set[T]) Intersect(s2 *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range s.data {
		if _, ok := s2.data[key]; ok {
			result.Add(key)
		}
	}
	return result
}

func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	result := NewSet[T]()
	for key := range s.data {
		result.Add(key)
	}
	for key := range s2.data {
		result.Add(key)
	}
	return result
}

func (s *Set[T]) ToList() []T {
	keys := make([]T, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}
	return keys
}

// -------------------------------

// 并发安全的 Set
type SafeSet[T comparable] struct {
	data *sync.Map
}

func NewSafeSet[T comparable]() *SafeSet[T] {
	return &SafeSet[T]{
		data: &sync.Map{},
	}
}

func (s *SafeSet[T]) Add(key ...T) {
	for _, key := range key {
		s.data.Store(key, struct{}{})
	}
}

func (s *SafeSet[T]) Del(key ...T) {
	for _, key := range key {
		s.data.Delete(key)
	}
}

func (s *SafeSet[T]) Exist(key T) bool {
	_, ok := s.data.Load(key)
	return ok
}

func (s *SafeSet[T]) Intersect(s2 *SafeSet[T]) *SafeSet[T] {
	result := NewSafeSet[T]()
	s.data.Range(func(key, _ interface{}) bool {
		if _, ok := s2.data.Load(key); ok {
			result.Add(key.(T))
		}
		return true
	})
	return result
}

func (s *SafeSet[T]) Union(s2 *SafeSet[T]) *SafeSet[T] {
	result := NewSafeSet[T]()
	s.data.Range(func(key, _ interface{}) bool {
		result.Add(key.(T))
		return true
	})
	s2.data.Range(func(key, _ interface{}) bool {
		result.Add(key.(T))
		return true
	})

	return result
}

func (s *SafeSet[T]) ToList() []T {
	keys := []T{}
	s.data.Range(func(key, _ interface{}) bool {
		keys = append(keys, key.(T))
		return true
	})
	return keys
}
