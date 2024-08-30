package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/stripe/stripe-go/v79/paymentintent"
)

func main() {
	stripe.Key = os.Getenv("SK_TEST")
	if stripe.Key == "" {
		log.Fatal("No stripe API key provided, please provide SK_TEST variable")
	}

	params := stripe.CustomerParams{
		Email:         stripe.String("example@example.com"),
		Name:          stripe.String("exampleName"),
		PaymentMethod: stripe.String("pm_card_visa"),
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String("pm_card_visa"),
		},
		Metadata: map[string]string{
			"timestamp": time.Now().String(),
		},
	}

	_, err := customer.New(&params)
	if err != nil {
		log.Fatal(err)
	}

	paramsPaymentIntent := stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1099),
		Currency: stripe.String(string(stripe.CurrencyPLN)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, err := paymentintent.New(&paramsPaymentIntent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
