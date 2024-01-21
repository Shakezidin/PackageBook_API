package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

func ViewPackage(ctx *gin.Context, client pb.UserClient) {
	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		fmt.Println("package id missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.UserViewPackage(ctxt, &pb.UserView{
		Id: int64(packageId),
	})
	if err != nil {
		log.Printf("activity fetching  error", err.Error())
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

func ViewCatagories(ctx *gin.Context,client pb.UserClient){
	var ctxt = context.Background()
	response, err := client.UserViewCatagories(ctxt, &pb.UserView{})

	if err != nil {
		log.Printf("catagories fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("catagories fetched succesfully"),
		"data":    response,
	})
}