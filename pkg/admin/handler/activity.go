package handler

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

// ViewActivity handles the HTTP request to view an activity by its ID.
func ViewActivity(ctx *gin.Context, client pb.AdminClient) {
	// Get the activity ID from the request header.
	activityIdStr := ctx.GetHeader("id")

	// Convert the activity ID to an integer.
	activityId, err := strconv.Atoi(activityIdStr)
	if err != nil {
		errMsg := "Invalid activity ID"
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  errMsg,
		})
		return
	}

	// Create a context for the gRPC request.
	ctxt := context.Background()

	// Call the gRPC service to fetch the activity by ID.
	response, err := client.AdminViewActivity(ctxt, &pb.AdminView{Id: int64(activityId)})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Check if the response is nil or the activity ID is 0, indicating that the activity was not found.
	if response == nil || response.ActivityId == 0 {
		errMsg := "Activity not found"
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"Status": http.StatusNotFound,
			"Error":  errMsg,
		})
		return
	}

	// Respond with the fetched activity.
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Activity fetched successfully",
		"Data":    response,
	})
}
