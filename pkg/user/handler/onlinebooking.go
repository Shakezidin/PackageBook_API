package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	pb "github.com/Shakezidin/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

func OnlinePayment(ctx *gin.Context, client pb.UserClient) {
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
		"UserId":           response.UserId,
		"TotalFare":        response.TotalFare,
		"BookingReference": response.BookingReference,
		"Email":            response.Email,
		"OrderId":          response.OrderId,
	})
}
