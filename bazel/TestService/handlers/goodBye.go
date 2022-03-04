package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func (g *GoodBye) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//TODO implement me
	res.Write([]byte("Byee"))
	return
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}
