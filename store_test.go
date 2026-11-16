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