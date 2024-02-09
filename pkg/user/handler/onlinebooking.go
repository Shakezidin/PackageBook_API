package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

func OnlinePayment(ctx *gin.Context, client pb.UserClient, typ string) {
	refId := ctx.Query("refid")
	if refId == "" {
		log.Println("reference id is empty")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("reference id is empty"),
		})
		return
	}
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	response, err := client.UserOnlinePayment(cont, &pb.UserBooking{
		Typ:   typ,
		RefId: refId,
	})

	if err != nil {
		log.Printf("payment unsuccesful err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusCreated, "app.html", gin.H{
		"userID":           response.UserId,
		"total":            response.TotalFare,
		"BookingReference": response.BookingReference,
		"Email":            response.Email,
		"orderID":          response.OrderId,
	})
}

func PaymentSuccess(ctx *gin.Context, client pb.UserClient) {
	paymentAmount := ctx.Query("total")
	signature := ctx.Query("signature")
	refID := ctx.Query("ref_id")
	orderID := ctx.Query("order_id")
	paymentId := ctx.DefaultQuery("payment_id", "")
	if refID == "" {
		log.Println("booking reference is not present")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("booking reference is empty"),
		})
		return
	}
	fmt.Println(orderID)
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	response, err := client.UserPaymentConfirmed(cont, &pb.UserPaymentConfirmedRequest{
		PaymentId:   paymentId,
		ReferenceID: refID,
		OrderID:     orderID,
		Signature:   signature,
		Total:       paymentAmount,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf(" Booking Confirmed"),
		"data":    response,
	})
}

func PaymentSuccessPage(ctx *gin.Context, client pb.UserClient) {
	ctx.HTML(http.StatusOK, "success.html", gin.H{
		"paymentID": ctx.Query("booking_reference"),
	})
}

// func handleStripeWebhook(c *gin.Context) {
// 	// Read the request body
// 	body, err := c.GetRawData()
// 	if err != nil {
// 		log.Printf("Error reading request body: %v", err)
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 		return
// 	}

// 	// Verify the Stripe webhook signature
// 	endpointSecret := "we_1OdCkCSCYyyrHDXu8erS4zLY"
// 	event, err := webhook.ConstructEvent(body, c.GetHeader("Stripe-Signature"), endpointSecret)
// 	if err != nil {
// 		log.Printf("Error verifying webhook signature: %v", err)
// 		c.AbortWithStatus(http.StatusBadRequest)
// 		return
// 	}

// 	// Handle the payment intent succeeded event
// 	if event.Type == "payment_intent.succeeded" {
// 		paymentIntent, ok := event.Data.Object["id"].(string)
// 		if !ok {
// 			log.Println("Payment intent ID not found in webhook data")
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}
// 		fmt.Println(paymentIntent)

// 		// Retrieve payment details from Stripe using the payment intent ID
// 		// Verify the payment details and save them into your database

// 		// Respond to the webhook with a 200 status code
// 		c.Status(http.StatusOK)
// 		return
// 	}

// 	// For other types of events, simply respond with a 200 status code
// 	c.Status(http.StatusOK)
// }
