package router

import (
	"encoding/json"
	"github.com/awesomebfm/go-store-backend/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/webhook"
	"net/http"
)

type CreateCheckoutBody struct {
	Items  []RequestItem `json:"items"`
	UserID string        `json:"user_id"`
}

type RequestItem struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

func HandleCreateCheckout(c *gin.Context) {
	var body CreateCheckoutBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Translate

	session, err := middleware.CreateSession()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"url": session.URL})
}

func HandleWebhook(c *gin.Context) {
	webhookSecret := ""
	payload, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	event, err := webhook.ConstructEvent(payload, c.Request.Header.Get("Stripe-Signature"), webhookSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Webhook signature verification failed"})
		return
	}

	switch event.Type {
	case "checkout.session.completed":
		var session stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &session)
		if err != nil {
			return
		}
	default:
		c.JSON(http.StatusOK, gin.H{"message": "Unhandled event type"})
	}
}
