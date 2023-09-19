package middleware

import (
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
)

func CreateSession() (*stripe.CheckoutSession, error) {
	stripe.Key = ""

	var orderItems []*stripe.CheckoutSessionLineItemParams

	params := &stripe.CheckoutSessionParams{
		LineItems:  orderItems,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String("https://google.com" + "?success=true"),
		CancelURL:  stripe.String("https://google.com" + "?canceled=true"),
	}

	s, err := session.New(params)

	if err != nil {
		return nil, err
	}
	return s, nil
}

func SaveOrder() {

}
