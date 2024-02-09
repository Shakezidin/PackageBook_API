package handler

import (
	"context"
	"net/http"

	"github.com/Shakezidin/middleware"
	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

func ValidateRefreshToken(ctx *gin.Context, client pb.UserClient) {
	id, err := middleware.ValidateRefreshToken(ctx, "user")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	var ctxt = context.Background()
	response, _ := client.UserRefreshToken(ctxt, &pb.TokenData{
		Role: "user",
		Id:   id,
	})
	ctx.JSON(200, gin.H{
		"status": http.StatusAccepted,
		"data":   response,
	})
}
