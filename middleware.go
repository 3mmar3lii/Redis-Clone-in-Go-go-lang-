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

func (l* LoggingMiddleware)NewLoggingMiddleware(inner store.Storer) *LoggingMiddleware{

	return &LoggingMiddleware{
		inner: l.inner,
		logger: log.New(os.Stdout,"[log] ",log.Ltime|log.Lmicroseconds),
	}
}