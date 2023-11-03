# shop_sample_go

0. Create a stripe test account

Get a publishable key from test dashboard

~~~
https://dashboard.stripe.com/test/apikeys
~~~

Paste it in frontend/src/StripePayment.jsx

~~~
const stripePromise = loadStripe(
  [your key]
)
~~~

??? Another secret key that right now in main.go but should be stored anywhere else ???

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

Choose an item and leave the shipping data empty.
After that you can fill the credit card data with test data

~~~
https://stripe.com/docs/testing
~~~

You will see completed payment in the dashboard.

~~~
https://dashboard.stripe.com/test/payments
~~~
