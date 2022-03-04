package handlers

import (
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func (g *Hello) ServeHTTP(res http.ResponseWriter, request *http.Request) {
	//TODO implement me
	res.Write([]byte("Hello World"))
	return
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
