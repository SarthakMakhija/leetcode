package lru_cache

import (
	"testing"
)

func TestLRUCacheGet1(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)

	value := lruCache.Get(1)
	if value != 1 {
		t.Fatalf("Expected 1 for a get by key = 1 received %v", value)
	}
}

func TestLRUCacheGet2(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)

	value := lruCache.Get(2)
	if value != 2 {
		t.Fatalf("Expected 2 for a get by key = 2 received %v", value)
	}
}

func TestLRUCacheGet3(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)

	value := lruCache.Get(1)
	if value != -1 {
		t.Fatalf("Expected -1 for a get by key = 1 received %v", value)
	}
}

func TestLRUCacheGet4(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)

	value := lruCache.Get(3)
	if value != 3 {
		t.Fatalf("Expected 3 for a get by key = 3 received %v", value)
	}
}

func TestLRUCacheGet5(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Get(1)
	lruCache.Put(3, 3)

	value := lruCache.Get(3)
	if value != 3 {
		t.Fatalf("Expected 3 for a get by key = 3 received %v", value)
	}
}

func TestLRUCacheGet6(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Get(1)
	lruCache.Put(3, 3)

	value := lruCache.Get(1)
	if value != 1 {
		t.Fatalf("Expected 1 for a get by key = 1 received %v", value)
	}
}

func TestLRUCacheGet7(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Get(1)
	lruCache.Put(3, 3)

	value := lruCache.Get(2)
	if value != -1 {
		t.Fatalf("Expected -1 for a get by key = 2 received %v", value)
	}
}

func TestLRUCacheGet8(t *testing.T) {
	lruCache := Constructor(1)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)

	value := lruCache.Get(2)
	if value != 2 {
		t.Fatalf("Expected 2 for a get by key = 2 received %v", value)
	}
}

func TestLRUCacheGet9(t *testing.T) {
	lruCache := Constructor(1)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)

	value := lruCache.Get(1)
	if value != -1 {
		t.Fatalf("Expected -1 for a get by key = 1 received %v", value)
	}
}

func TestLRUCacheGet10(t *testing.T) {
	lruCache := Constructor(1)
	lruCache.Get(6)
	lruCache.Get(8)
	lruCache.Put(12, 1)
	lruCache.Get(2)
	lruCache.Put(15, 11)
	lruCache.Put(5, 2)
	lruCache.Put(1, 15)
	lruCache.Put(4, 2)
	lruCache.Get(5)
	lruCache.Put(15, 15)

	value := lruCache.Get(15)
	if value != 15 {
		t.Fatalf("Expected 15 for a get by key = 15 received %v", value)
	}
}
