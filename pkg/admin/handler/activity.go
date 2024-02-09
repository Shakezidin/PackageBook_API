package handler

import (
	"context"
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
		errMsg := "invalid activity ID"
		log.Println(errMsg)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errMsg,
		})
		return
	}

	ctxt := context.Background()
	response, err := client.AdminViewActivity(ctxt, &pb.AdminView{Id: int64(activityId)})
	if err != nil {
		errMsg := "error fetching activity"
		log.Printf("%s: %v", errMsg, err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errMsg,
		})
		return
	}

	if response == nil || response.ActivityId == 0 {
		errMsg := "activity not found"
		log.Println(errMsg)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errMsg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "activity fetched successfully",
		"data":    response,
	})
}
