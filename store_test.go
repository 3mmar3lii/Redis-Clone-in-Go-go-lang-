package main

import (
	"reflect"
	"testing"
)

func TestKeys_ReturnAllKeysSorted(t *testing.T){

	store := NewStore()
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
	store :=NewStore()
	if len(store.Keys()) > 0 {
		t.Errorf("Store Keys() = Is not emtpy %v",store.Keys())
	}
}

func TestSetGet_RoundTrip(t *testing.T){
	store :=NewStore()
	store.Set("hello","world")
	if val,ok :=store.Get("hello"); !ok || val!="world"{
		t.Errorf("This key is not exist ")
	} 
	if val,ok := store.Get("missing"); ok || val !=""{
		t.Error("This key is not exist ")
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
			store := NewStore()
			for k, v := range tc.setup {
				store.Set(k, v)
			}

			store.Delete(tc.deleteKey)

			if got := store.Len(); got != tc.wantLen {
				t.Errorf("after Delete(%q): Len() = %d, want %d", tc.deleteKey, got, tc.wantLen)
			}

			if _, ok := store.Get(tc.deleteKey); ok {
				t.Errorf("after Delete(%q): key still present", tc.deleteKey)
			}
		})
	}
}
