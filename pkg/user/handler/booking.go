package handler

import (
	"errors"
	"log"
	"net/http"

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
	Travellers []map[string]string `json:"travellers"`
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
		log.Println("pacakge id required")
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

	userID, err := extractUserID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errors.New("error getting value from token"),
		})
		return
	}

	for _, travellerMap := range travellerDetails.Travellers {
		activityIDs := []string{travellerMap["id"]}
		td = append(td, &pb.UserTravellerDetails{
			Name:       travellerMap["name"],
			Age:        travellerMap["age"],
			Gender:     travellerMap["gender"],
			ActivityId: activityIDs,
		})
	}

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
		"message": "successful in retrieving search details",
		"data":    response,
	})
}
