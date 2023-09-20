package service

import (
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/checkout/session"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
)

type CheckoutItemDto struct {
	PriceID  string `json:"price_id"`
	Quantity int    `json:"quantity"`
}

func GenerateCheckoutSession(items []CheckoutItemDto, customerId primitive.ObjectID) (string, error) {
	stripe.Key = os.Getenv("STRIPE_PRIVATE_KEY")

	var sessionProducts []*stripe.CheckoutSessionLineItemParams
	for _, item := range items {
		sessionProduct := &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.PriceID),
			Quantity: stripe.Int64(int64(item.Quantity)),
		}
		sessionProducts = append(sessionProducts, sessionProduct)
	}

	params := &stripe.CheckoutSessionParams{
		LineItems:         sessionProducts,
		Mode:              stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:        stripe.String("https://google.com" + "?success=true"),
		CancelURL:         stripe.String("https://google.com" + "?canceled=true"),
		ClientReferenceID: stripe.String(customerId.Hex()),
	}

	s, err := session.New(params)

	if err != nil {
		return "", err
	}
	return s.URL, nil
}
