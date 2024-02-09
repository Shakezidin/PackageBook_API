package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

func ViewUser(ctx *gin.Context, client pb.AdminClient) {
	idStr := ctx.GetHeader("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err,
		})
	}

	ctxt := context.Background()
	response, err := client.AdminViewUser(ctxt, &pb.AdminView{
		Id: int64(id),
	})

	if err != nil {
		errMsg := "error while fetching user"
		log.Printf("%s: %v", errMsg, err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errMsg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user fetched successfully",
		"data":    response,
	})
}

func ViewUsers(ctx *gin.Context, client pb.AdminClient) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	var ctxt = context.Background()
	response, err := client.AdminViewUsers(ctxt, &pb.AdminView{
		Page: int64(page),
	})

	if err != nil {
		log.Printf("error while users packages: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "failed to fetch users",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "users fetched successfully",
		"data":    response,
	})
}
