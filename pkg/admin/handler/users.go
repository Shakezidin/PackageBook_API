package handler

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

// ViewUser handles the HTTP request to view a specific user.
func ViewUser(ctx *gin.Context, client pb.AdminClient) {
	idStr := ctx.GetHeader("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Invalid user ID",
		})
		return
	}

	// Create a context
	ctxt := context.Background()

	// Call the gRPC service to fetch the user
	response, err := client.AdminViewUser(ctxt, &pb.AdminView{
		ID: int64(id),
	})

	// Handle errors
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Respond with fetched user
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "User fetched successfully",
		"Data":    response,
	})
}

// ViewUsers handles the HTTP request to view a list of users.
func ViewUsers(ctx *gin.Context, client pb.AdminClient) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	// Create a context
	ctxt := context.Background()

	// Call the gRPC service to fetch users
	response, err := client.AdminViewUsers(ctxt, &pb.AdminView{
		Page: int64(page),
	})

	// Handle errors
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Respond with fetched users
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Users fetched successfully",
		"Data":    response,
	})
}
