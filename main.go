package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/steven", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("I am the mans."))
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
