package ttl

import (
	"testing"
	"time"
)

func TestTTLStore_GetAfterExpiry(t *testing.T) {
	s := NewTTLStore(time.Second * 3)
	_ = s.Set("k", "v")

	// simulate waiting until key expires
	time.Sleep(5 * time.Second)

	if val, err := s.Get("k"); err == nil || val != "" {
		t.Errorf("expected error after TTL, got nil")
	}
}

func TestTTLStore_ExpiredKeyEvicted(t *testing.T) {
	store := NewTTLStore(time.Microsecond * 50)
	_ = store.Set("k", "v")

	// simulate waiting until key expires
	time.Sleep(100 * time.Millisecond)

	val, err := store.Get("k")
	if err == nil {
		t.Errorf("expected an error for expired key 'k', got value %q and nil error", val)
	}
	if val != "" {
		t.Errorf("expected empty string for expired key, got %q", val)
	}
}


