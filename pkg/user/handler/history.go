package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Shakezidin/middleware"
	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

// ViewHistory retrieves user's booking history.
func ViewHistory(ctx *gin.Context, client pb.UserClient) {
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)

	// Extract user ID from token
	_, id, err := middleware.ValidateToken(ctx, "user")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.UserViewHistory(ctxt, &pb.UserView{
		Page: int64(pageInt),
		Id:   int64(userId),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "history fetched successfully",
		"data":    response,
	})
}

// ViewBooking retrieves user's booking details.
func ViewBooking(ctx *gin.Context, client pb.UserClient) {
	var ctxt = context.Background()
	id := ctx.GetHeader("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Booking ID missing",
		})
		return
	}
	response, err := client.UserViewBooking(ctxt, &pb.UserView{
		Id: int64(Id),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
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

// PackageCancel cancels a user's booking.
func PackageCancel(ctx *gin.Context, client pb.UserClient) {
	id := ctx.GetHeader("id")
	startdate := ctx.GetHeader("startdate")
	Id, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Booking ID missing",
		})
		return
	}

	strtdate, err := time.Parse("02-01-2006", startdate)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Error while parsing date",
		})
		return
	}

	duration := strtdate.Sub(time.Now())
	if duration <= 24*time.Hour {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Cancellation time out",
		})
		return
	}

	// Extract user ID from token
	_, userID, err := middleware.ValidateToken(ctx, "user")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Status": http.StatusUnauthorized,
			"Error":  err.Error(),
		})
		return
	}
	userId, _ := strconv.Atoi(userID)

	var ctxt = context.Background()
	response, err := client.UserCancelBooking(ctxt, &pb.UserView{
		Id:     int64(Id),
		UserId: int64(userId),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Package cancelled successfully",
		"Data":    response,
	})
}
