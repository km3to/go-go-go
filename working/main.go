package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oooops!", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "\nHello, %s\n\n", d)
	})

	http.HandleFunc("/bye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("BoodBye World!")
		d, _ := ioutil.ReadAll(r.Body)
		log.Printf("Data: %s\n", d)
	})

	http.ListenAndServe(":9090", nil)
}
