package store

import "sync"

type Store struct {
	sync.RWMutex
	m map[string]string
}

func New() *Store {
	return &Store{
		m: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	v, ok := s.m[key]
	return v, ok
}

func (s *Store) Put(key, value string) {
	s.Lock()
	defer s.Unlock()
	s.m[key] = value
}

func (s *Store) Delete(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, key)
}
