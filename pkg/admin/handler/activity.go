package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

func ViewActivity(ctx *gin.Context, client pb.AdminClient) {
	activityIdStr := ctx.GetHeader("id")
	activityId, err := strconv.Atoi(activityIdStr)
	if err != nil {
		fmt.Println("destinationID missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.AdminViewActivity(ctxt, &pb.AdminView{
		Id: int64(activityId),
	})

	if err != nil {
		log.Printf("package fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("activity fetched succesfully"),
		"data":    response,
	})
}
