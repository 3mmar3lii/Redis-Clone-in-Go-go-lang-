package main

import "fmt"

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string,bool){
	val,ok := s.data[key]
	return val,ok
}


func (s *Store) Set(key,value string){
	s.data[key]=value
}


func (s * Store) Delete(key string) {
	if _,ok := s.data[key]; !ok {
		return
	}
	delete(s.data,key)
}

func NewStore() *Store{
	return &Store{
		data :make(map[string]string),
	}
}

func main(){
	store := NewStore()
	store.Set("key1","val1")

	val,_ := store.Get("key1")
	fmt.Println(val)
	//fmt.Println("ok",ok)
}