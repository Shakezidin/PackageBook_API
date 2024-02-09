package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

func ViewCoordinators(ctx *gin.Context, client pb.AdminClient) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	status := ctx.DefaultQuery("status", "true")

	ctxt := context.Background()
	response, err := client.AdminViewCoordinators(ctxt, &pb.AdminView{
		Status: status,
		Page:   int64(page),
	})

	if err != nil {
		errMsg := "error while fetching coordinators"
		log.Printf("%s: %v", errMsg, err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errMsg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "coordinators fetched successfully",
		"data":    response,
	})
}
