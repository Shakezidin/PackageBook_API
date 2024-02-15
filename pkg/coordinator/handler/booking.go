package handler

import (
	"context"
	"net/http"
	"strconv"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
)

// ViewBookings retrieves all bookings related to a package.
func ViewBookings(ctx *gin.Context, client cpb.CoordinatorClient) {
	// Retrieve package ID from header
	id := ctx.GetHeader("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Package ID missing or invalid",
		})
		return
	}

	// Create a background context
	var ctxt = context.Background()
	// Call gRPC service to view booking history
	response, err := client.ViewHistory(ctxt, &cpb.View{
		Status: "true",
		Id:     int64(ID),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": "booking fetched succesfully",
		"data":    response,
	})
}

// ViewBooking retrieves details of a specific booking.
func ViewBooking(ctx *gin.Context, client cpb.CoordinatorClient) {
	// Retrieve booking ID from header
	id := ctx.GetHeader("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Booking ID missing or invalid",
		})
		return
	}

	// Create a background context
	var ctxt = context.Background()
	// Call gRPC service to view booking details
	response, err := client.ViewBooking(ctxt, &cpb.View{
		Id: int64(ID),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Booking fetched successfully",
		"Data":    response,
	})
}

// ViewTraveller retrieves details of a specific traveller.
func ViewTraveller(ctx *gin.Context, client cpb.CoordinatorClient) {
	// Retrieve traveller ID from header
	id := ctx.GetHeader("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Traveller ID missing or invalid",
		})
		return
	}

	// Create a background context
	var ctxt = context.Background()
	// Call gRPC service to view traveller details
	response, err := client.ViewTraveller(ctxt, &cpb.View{
		Id: int64(ID),
	})

	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Traveller fetched successfully",
		"Data":    response,
	})
}
