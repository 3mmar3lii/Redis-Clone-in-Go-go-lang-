package main

import (
	"Redis-Clone/store"
	"log"
	"os"
)


type LoggingMiddleware struct {
	inner store.Storer
	logger  *log.Logger
}

func NewLoggingMiddleware(inner store.Storer) *LoggingMiddleware{

	return &LoggingMiddleware{
		inner: inner,
		logger: log.New(os.Stdout,"[log] ",log.Ltime|log.Lmicroseconds),
	}
}

func (l *LoggingMiddleware) Set(key,value string) error {
	l.logger.Printf("SET %q",key) // side efffect
	return l.inner.Set(key,value)
}

func (l *LoggingMiddleware) Get(key string) (string,error) {
	l.logger.Printf("GET %q",key)
	
	got,err := l.inner.Get(key)
	if err != nil {
		l.logger.Print("Error GET %q -> miss (%v) ",key,err)
	}
	l.logger.Printf("GET %q -> hit",key)
	return got,err
}

func (l *LoggingMiddleware)Delete(key string) error {
	l.logger.Printf("DELETE %q",key)
	return l.inner.Delete(key)
}

func (l *LoggingMiddleware)Len() int {
	got := l.inner.Len()
	l.logger.Printf("LEN %d",got) // side efffect
	return got
}

func (l *LoggingMiddleware)Keys() []string {
	got := l.inner.Keys()
	l.logger.Printf("KEYS %v",got) // side efffect
	return got
}