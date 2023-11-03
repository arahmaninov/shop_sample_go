# shop_sample_go

## About

This is a sample shop with backend made in Go. Payment processing api provided by Stripe.

![shop_sample_go](https://github.com/arahmaninov/shop_sample_go/assets/95240106/cb99dc5a-6c5f-418b-a1ae-1964a89e280e)

## Installation

0. Create a stripe test account
   (https://stripe.com/)

Get a publishable key from the dashboard

(https://dashboard.stripe.com/test/apikeys)

Paste it in frontend/src/StripePayment.jsx

~~~
const stripePromise = loadStripe(
  [your key]
)
~~~

Add a secret key from the dashboard to your environment (see .env.example).

1. Build and run the frontend client

~~~
npm install
~~~

~~~
npm start
~~~

( export Node options in case of errors: )

~~~
export NODE_OPTIONS=--openssl-legacy-provider
~~~

2. Run backend server

~~~
go run main.go
~~~

Choose an item, you can leave the shipping data empty.
After that you can fill the credit card data with test data

(https://stripe.com/docs/testing)

You will see completed payment in the dashboard.

(https://dashboard.stripe.com/test/payments)

![shop_sample_go2](https://github.com/arahmaninov/shop_sample_go/assets/95240106/8dab48ea-f4ce-407f-97a6-5b4f6d390bbd)

