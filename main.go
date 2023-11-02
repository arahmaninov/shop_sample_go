package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	
	err := http.ListenAndServe("localhost:4242", nil) // run the server
	if err != nil {
		log.Fatal(err)
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT CALLED: handleAbout")
	
	response := []byte("Test server built in Go")
	w.Write(response)
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT CALLED: handleCreatePaymentIntent")
}

