package main

import "fmt"

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string,bool){
	return "" ,false
}


func (s *Store) Set(key,value string){
	//return 
}


func NewStore() *Store{
	return &Store{
		data :make(map[string]string),
	}
}

func main(){
	store := NewStore()
	store.Get("key1")
	store.Set("key1","val1")
	fmt.Println("")
}