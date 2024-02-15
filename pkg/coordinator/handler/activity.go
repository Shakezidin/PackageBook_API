package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AddActivity handles the addition of a new activity by the coordinator.
func AddActivity(ctx *gin.Context, client cpb.CoordinatorClient) {
	var activity dto.AddActivities

	// Retrieve destination ID from header
	destinationIDStr := ctx.GetHeader("id")
	destinationID, err := strconv.Atoi(destinationIDStr)
	if err != nil {
		// Return error response if destination ID is missing or invalid
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Destination ID missing or invalid",
		})
		return
	}

	// Bind JSON data to activity struct
	if err := ctx.BindJSON(&activity); err != nil {

		// Return error response if JSON data is invalid
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Invalid JSON data",
		})
		return
	}

	// Validate the activity struct
	validate := validator.New()
	err = validate.Struct(activity)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		for _, e := range err.(validator.ValidationErrors) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("error in field %v, error: %v", e.Field(), e.Tag()),
			})
		}
		return
	}

	// Create a background context
	ctxt := context.Background()
	// Call gRPC service to add activity
	response, err := client.CoordinatorAddActivity(ctxt, &cpb.Activity{
		Activityname:  activity.ActivityName,
		Description:   activity.Description,
		Location:      activity.Location,
		ActivityType:  activity.ActivityType,
		Amount:        int64(activity.Price),
		Date:          activity.Date,
		Time:          activity.Time,
		DestinationId: int64(destinationID),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	// Return success response with created activity data
	ctx.JSON(http.StatusCreated, gin.H{
		"Status": http.StatusCreated,
		"Data":   response,
	})
}

// ViewActivity fetches a specific activity by its ID.
func ViewActivity(ctx *gin.Context, client cpb.CoordinatorClient) {
	// Retrieve activity ID from header
	activityIDStr := ctx.GetHeader("id")
	activityID, err := strconv.Atoi(activityIDStr)
	if err != nil {
		// Return error response if activity ID is missing or invalid
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Activity ID missing or invalid",
		})
		return
	}

	// Create a background context
	ctxt := context.Background()
	// Call gRPC service to fetch activity by ID
	response, err := client.CoordinatorViewActivity(ctxt, &cpb.View{
		Id: int64(activityID),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	// Return success response with fetched activity data
	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": "activity fetched succesfully",
		"data":    response,
	})
}
