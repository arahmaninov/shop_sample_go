package main

import (
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	//"github.com/stripe/stripe-go/v76/customer"
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
	
	response := []byte("Test server built in Go")
	_, err := w.Write(response)
	if err != nil {
		log.Println(err)
	}
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}


	var request struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Address1 string `json:"address_1"`
		Address2 string `json:"address_2"`
		City string `json:"city"`
		State string `json:"state"`
		Zip string `json:"zip"`
		Country string `json:"country"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Stripe api 
	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(calculateOrderAmount(request.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	paymentIntent, err := paymentintent.New(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(paymentIntent.ClientSecret)
}

func calculateOrderAmount(productId string) int64 {
	switch productId {
	case "Forever Pants":
		return 26000
	case "ForeverShirt":
		return 15500
	case "Forever Shorts":
		return 30000
	}

	return 0
}
