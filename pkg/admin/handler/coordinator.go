package handler

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

// ViewCoordinators handles the HTTP request to view coordinators.
func ViewCoordinators(ctx *gin.Context, client pb.AdminClient) {
	// Get query parameters
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	status := ctx.DefaultQuery("status", "true")

	// Create a context
	ctxt := context.Background()

	// Call the gRPC service to fetch coordinators
	response, err := client.AdminViewCoordinators(ctxt, &pb.AdminView{
		Status: status,
		Page:   int64(page),
	})

	// Handle errors
	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Respond with fetched coordinators
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Coordinators fetched successfully",
		"Data":    response,
	})
}
