package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(logger *log.Logger) *Goodbye {
	return &Goodbye{logger}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goodbye!"))
}
