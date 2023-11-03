package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	
	stripeKey, exists := os.LookupEnv("STRIPE_KEY")
	if !exists { 
		log.Print("There is no Stripe key in .env file") 
	}

	stripe.Key = stripeKey

	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	
	err := http.ListenAndServe(":4242", nil) // run the server
	if err != nil {
		log.Fatal(err)
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	
	response := []byte("Test shop server working with Stripe API")
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


	var response struct {
		ClientSecret string `json:"clientSecret"`
	}

	response.ClientSecret = paymentIntent.ClientSecret

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = io.Copy(w, &buf)
	if err != nil {
		fmt.Println(err)
	}
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
