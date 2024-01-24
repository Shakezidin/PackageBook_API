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

func ViewDestination(ctx *gin.Context, client pb.AdminClient) {
	destinationIdStr := ctx.GetHeader("id")
	destinationId, err := strconv.Atoi(destinationIdStr)
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
	response, err := client.AdminViewDestination(ctxt, &pb.AdminView{
		Id: int64(destinationId),
	})

	if err != nil {
		log.Printf("error while fetching destination", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("destination fetched succesfully"),
		"data":    response,
	})
}
