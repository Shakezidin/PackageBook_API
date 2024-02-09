package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

func ViewDestination(ctx *gin.Context, client pb.AdminClient) {
	destinationIdStr := ctx.GetHeader("id") // Retrieve destination ID from path parameter
	destinationId, err := strconv.Atoi(destinationIdStr)
	if err != nil {
		errMsg := "destination ID missing or invalid"
		log.Printf("%s: %v", errMsg, err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errMsg,
		})
		return
	}

	ctxt := context.Background()
	response, err := client.AdminViewDestination(ctxt, &pb.AdminView{
		Id: int64(destinationId),
	})

	if err != nil {
		errMsg := "error while fetching destination"
		log.Printf("%s: %v", errMsg, err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errMsg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "destination fetched successfully",
		"data":    response,
	})
}
