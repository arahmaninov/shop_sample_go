package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/health")
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	
	errLaS := http.ListenAndServe("localhost:4242", nil) // run the server
	if errLaS != nil {
		log.Fatal(errLaS)
	}
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT CALLED: handleCreatePaymentIntent")
}
