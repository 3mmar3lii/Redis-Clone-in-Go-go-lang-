package main

import (
	"Redis-Clone/store"
	"fmt"
	"time"
)

type MetrixMiddleware struct {
	inner store.Storer
	setCalls int
	getCalls int
	deleteCalls int
	lenCalls int
	getMisses int
	totalSetLatency time.Duration
	totalGetLatency time.Duration
}


func NewMatrixMiddleware(store store.Storer) * MetrixMiddleware {
	return &MetrixMiddleware{
		inner : store,
	}
}


func (m *MetrixMiddleware)Set(key,value string) error{
	now := time.Now()
	m.setCalls++
	m.totalSetLatency += time.Since(now)
	return m.inner.Set(key,value)
}


func (m *MetrixMiddleware)Get(key,value string) (string,error){
	now := time.Now()
	val,err := m.inner.Get(key)
	m.getCalls++
	if err != nil {
		m.getMisses++
	}
	m.totalGetLatency += time.Since(now)
	return val,err
}


func (m *MetrixMiddleware)Delete(key string) error{
	m.deleteCalls++
	return m.inner.Delete(key)
}

func (m *MetrixMiddleware)Len() int{
	m.lenCalls++
	return m.inner.Len()
}


func (m *MetrixMiddleware)Keys() []string{
	return m.inner.Keys()
}

func (m *MetrixMiddleware) Report() {
	var avgGet, avgSet time.Duration
	if m.getCalls > 0 {
		avgGet = m.totalGetLatency / time.Duration(m.getCalls)
	}
	if m.setCalls > 0 {
		avgSet = m.totalSetLatency / time.Duration(m.setCalls)
	}

	fmt.Printf(
		"=== Metrics Report ===\n"+
			"GET Calls:    %d (Hits: %d | Misses: %d)\n"+
			"SET Calls:    %d\n"+
			"DELETE Calls: %d\n"+
			"LEN Calls:    %d\n"+
			"Avg GET:      %v\n"+
			"Avg SET:      %v\n",
		m.getCalls, m.getCalls-m.getMisses, m.getMisses,
		m.setCalls,
		m.deleteCalls,
		m.lenCalls,
		avgGet,
		avgSet,
	)
}
