package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Shakezidin/middleware"
	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

// TravellerDetail represents details of a traveller
type TravellerDetail struct {
	Name       string   `json:"name"`
	Age        string   `json:"age"`
	Gender     string   `json:"gender"`
	ActivityId []string `json:"id"`
}

type TravellerDetails struct {
	Travellers []TravellerDetail `json:"travellers"`
}

func AddTraveller(ctx *gin.Context, client pb.UserClient) {
	pkgId := ctx.GetHeader("id")
	if pkgId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  "package id missing",
			"status": http.StatusBadRequest,
		})
		return
	}
	var travellerDetails TravellerDetails
	var td []*pb.UserTravellerDetails

	if err := ctx.ShouldBindJSON(&travellerDetails); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	email, userID, err := middleware.ValidateToken(ctx, "user")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("error getting value from token"),
		})
		return
	}

	for _, travellerMap := range travellerDetails.Travellers {
		activityIDs := travellerMap.ActivityId
		td = append(td, &pb.UserTravellerDetails{
			Name:       travellerMap.Name,
			Age:        travellerMap.Age,
			Gender:     travellerMap.Gender,
			ActivityId: activityIDs,
		})
	}

	ctx.Set("registered_email", email)
	response, err := client.UserTravellerDetails(ctx, &pb.UserTravellerRequest{
		TravellerDetails: td,
		UserId:           userID,
		PackageId:        pkgId,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "user details saved successfully",
		"data":    response,
	})
}

// AdvancePayment handles the advance payment for a booking.
func AdvancePayment(ctx *gin.Context, client pb.UserClient) {
	refID := ctx.GetHeader("refid")
	if refID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("reference ID is empty"),
		})
		return
	}
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	email, userID, err := middleware.ValidateToken(ctx, "user")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("email id not present in jwt token, please login again"),
		})
		return
	}

	userId, _ := strconv.Atoi(userID)

	response, err := client.UserOfflineBooking(cont, &pb.UserBooking{
		RefId:  refID,
		UserId: int64(userId),
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Status":  http.StatusCreated,
		"Message": fmt.Sprintf("%v offline booking success", email),
		"Data":    response,
	})
}
