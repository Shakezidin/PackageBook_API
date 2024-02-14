package handler

import (
	"context"
	"net/http"

	"github.com/Shakezidin/middleware"
	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

// ValidateRefreshToken validates the refresh token and generates a new access token.
func ValidateRefreshToken(ctx *gin.Context, client pb.UserClient) {
	id, err := middleware.ValidateRefreshToken(ctx, "user")
	if err != nil {
		handleError(ctx, err, "Error validating refresh token")
		return
	}

	ctxt := context.Background()
	response, err := client.UserRefreshToken(ctxt, &pb.TokenData{
		Role: "user",
		Id:   id,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status": http.StatusOK,
		"Data":   response,
	})
}
