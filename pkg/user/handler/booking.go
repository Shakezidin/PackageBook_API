package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Shakezidin/middleware"
	pb "github.com/Shakezidin/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

type TravellerDetail struct {
	Name       string   `json:"name"`
	Age        string   `json:"age"`
	Gender     string   `json:"gender"`
	ActivityId []string `json:"id"`
}

type TravellerDetails struct {
	Travellers []TravellerDetail `json:"travellers"`
}

func extractUserID(ctx *gin.Context) (string, error) {
	_, userID, err := middleware.ValidateToken(ctx, "user")
	if err != nil {
		return "", errors.New("error extracting user ID from token")
	}
	return userID, nil
}

func AddTraveller(ctx *gin.Context, client pb.UserClient) {
	pkgId := ctx.GetHeader("packageId")
	if pkgId == "" {
		log.Println("package id required")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  "package id missing",
			"status": http.StatusBadRequest,
		})
		return
	}
	var travellerDetails TravellerDetails
	var td []*pb.UserTravellerDetails

	if err := ctx.ShouldBindJSON(&travellerDetails); err != nil {
		log.Println("unable to bind JSON, err:", err.Error())
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
		log.Println("unable to bind JSON, err:", err.Error())
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

func AdvancePayment(ctx *gin.Context, client pb.UserClient) {
	refId := ctx.GetHeader("refid")
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

	email, userId, err := middleware.ValidateToken(ctx, "user")
	if err != nil {
		log.Println("user id not present in jwt token, please login again")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("email id not present in jwt token, please login again"),
		})
		return
	}

	userID, _ := strconv.Atoi(userId)

	response, err := client.UserOfflineBooking(cont, &pb.UserBooking{
		RefId:  refId,
		UserId: int64(userID),
	})

	if err != nil {
		log.Printf("unable to do payment for %v err: %v", email, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v offline booking success", email),
		"data":    response,
	})
}
