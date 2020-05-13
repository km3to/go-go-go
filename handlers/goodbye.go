package handlers

import (
	"log"
	"net/http"
)

// Goodbye handler
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye - constructor with DI for Goodbye
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byeeee\n"))
}
