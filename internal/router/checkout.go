package router

import (
	"encoding/json"
	"github.com/awesomebfm/go-store-backend/internal/database"
	"github.com/awesomebfm/go-store-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/webhook"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
)

type CreateCheckoutBody struct {
	Items  []RequestItem `json:"items"`
	UserID string        `json:"user_id"`
}

type RequestItem struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

func CreateCheckoutSession(c *gin.Context) {
	var body CreateCheckoutBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var items []service.CheckoutItemDto
	for i := 0; i < len(body.Items); i++ {
		productId, err := primitive.ObjectIDFromHex(body.Items[i].ID)
		if err != nil {
			continue
		}

		product, err := database.GetProductByID(productId)
		if err != nil {
			continue
		}

		if body.Items[i].Quantity > product.Inventory {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product is out of stock, or too many units requested!"})
			return
		}

		items = append(items, service.CheckoutItemDto{PriceID: product.StripeID, Quantity: body.Items[i].Quantity})
	}

	if len(items) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No items were included!"})
		return
	}

	if body.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No customer was provided!"})
		return
	}

	userId, err := primitive.ObjectIDFromHex(body.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while parsing user_id!"})
		return
	}

	url, err := service.GenerateCheckoutSession(items, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while generating checkout url!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}

func HandleWebhook(c *gin.Context) {
	webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
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
