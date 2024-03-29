package handler

import (
	"context"
	"net/http"
	"time"

	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

// OnlinePayment handles the online payment process.
func OnlinePayment(ctx *gin.Context, client pb.UserClient, typ string) {
	refID := ctx.Query("refid")
	if refID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Reference ID is empty",
		})
		return
	}
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	response, err := client.UserOnlinePayment(cont, &pb.UserBooking{
		Type:   typ,
		Ref_ID: refID,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusCreated, "app.html", gin.H{
		"userID":           response.User_ID,
		"total":            response.Total_Fare,
		"BookingReference": response.Booking_Reference,
		"Email":            response.Email,
		"orderID":          response.Order_ID,
	})
}

// PaymentSuccess handles the payment success confirmation.
func PaymentSuccess(ctx *gin.Context, client pb.UserClient) {
	paymentAmount := ctx.Query("total")
	signature := ctx.Query("signature")
	refID := ctx.Query("ref_id")
	orderID := ctx.Query("order_id")
	paymentID := ctx.DefaultQuery("payment_id", "")

	if refID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Booking reference is empty",
		})
		return
	}

	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	response, err := client.UserPaymentConfirmed(cont, &pb.UserPaymentConfirmedRequest{
		Payment_ID:   paymentID,
		Reference_ID: refID,
		Order_ID:     orderID,
		Signature:   signature,
		Total:       paymentAmount,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "Booking Confirmed",
		"data":    response,
	})
}

// PaymentSuccessPage returns the success page after payment.
func PaymentSuccessPage(ctx *gin.Context, client pb.UserClient) {
	ctx.HTML(http.StatusOK, "success.html", gin.H{
		"paymentID": ctx.Query("booking_reference"),
	})
}
