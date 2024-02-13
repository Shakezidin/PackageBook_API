package handler

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

// ViewDestination handles the HTTP request to view a destination.
func ViewDestination(ctx *gin.Context, client pb.AdminClient) {
	// Retrieve destination ID from header
	destinationIdStr := ctx.GetHeader("id")
	destinationId, err := strconv.Atoi(destinationIdStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Destination ID missing or invalid",
		})
		return
	}

	// Create a context
	ctxt := context.Background()

	// Call the gRPC service to fetch the destination
	response, err := client.AdminViewDestination(ctxt, &pb.AdminView{
		Id: int64(destinationId),
	})

	// Handle errors
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Respond with fetched destination
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Destination fetched successfully",
		"Data":    response,
	})
}
