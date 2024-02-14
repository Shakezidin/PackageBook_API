package handler

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

// ViewDestination retrieves destination details.
func ViewDestination(ctx *gin.Context, client pb.UserClient) {
	destinationIDStr := ctx.GetHeader("id")
	destinationID, err := strconv.Atoi(destinationIDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"Error":  "Destination ID missing",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.UserViewDestination(ctxt, &pb.UserView{
		Id: int64(destinationID),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Destination fetched successfully",
		"Data":    response,
	})
}

// ViewActivity retrieves activity details.
func ViewActivity(ctx *gin.Context, client pb.UserClient) {
	activityIDStr := ctx.GetHeader("id")
	activityID, err := strconv.Atoi(activityIDStr)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"Error":  "activity ID missing",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.UserViewActivity(ctxt, &pb.UserView{
		Id: int64(activityID),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Activity fetched successfully",
		"Data":    response,
	})
}
