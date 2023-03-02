package store

import "testing"

func TestPut(t *testing.T) {
	s := New()
	s.Put("a", "1")
	if v, ok := s.Get("a"); !ok || v != "1" {
		t.Fatal("expected a=1")
	}
}

func TestDelete(t *testing.T) {
	s := New()
	s.Put("a", "1")
	s.Delete("a")
	if _, ok := s.Get("a"); ok {
		t.Fatal("expected a to be deleted")
	}
}

func TestGet(t *testing.T) {
	s := New()
	s.Put("a", "1")
	if v, ok := s.Get("b"); ok || v != "" {
		t.Fatal("expected b to not exist")
	}
}
