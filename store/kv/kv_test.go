package kv

import (
	"reflect"
	"testing"
)

func TestKeys_ReturnAllKeysSorted(t *testing.T){

	store := NewStore(3)
	store.Set("ammar","3")
	store.Set("mai","344")
	store.Set("waleed","34")

	got := store.Keys()
	want := []string{"ammar","mai","waleed"}

	if !reflect.DeepEqual(got,want){
		t.Errorf("Keys () = %v, want %v",got,want)
	}
}


func TestKey_EmptyKeys(t *testing.T){
	store :=NewStore(3)
	if len(store.Keys()) > 0 {
		t.Errorf("Store Keys() = Is not emtpy %v",store.Keys())
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore(3)
	store.Set("hello", "world")

	// 1. Test retrieving a missing key (we EXPECT an error or empty value here)
	if val, err := store.Get("missing"); err == nil || val != "" {
		t.Errorf("expected error for missing key, got value: %q", val)
	}

	if val, err := store.Get("hello"); err != nil || val != "world" {
		t.Errorf("expected 'world', got %q (err: %v)", val, err)
	}
}



// Table driven test
func TestDelete(t *testing.T) {
	tests := []struct {
		name      string
		setup     map[string]string
		deleteKey string
		wantLen   int
	}{
		{"deletes existing key", map[string]string{"a": "1", "b": "2"}, "a", 1},
		{"no-op on missing key", map[string]string{"a": "1"}, "x", 1},
		{"no-op on empty store", map[string]string{}, "x", 0},
		{"deletes last remaining key", map[string]string{"only": "value"}, "only", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := NewStore(3)
			for k, v := range tc.setup {
				store.Set(k, v)
			}

			store.Delete(tc.deleteKey)

			if got := store.Len(); got != tc.wantLen {
				t.Errorf("after Delete(%q): Len() = %d, want %d", tc.deleteKey, got, tc.wantLen)
			}

			if _, err := store.Get(tc.deleteKey); err == nil {
				t.Errorf("after Delete(%q): key still present", tc.deleteKey)
			}
		})
	}
}
