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

func ViewPackages(ctx *gin.Context, client pb.AdminClient) {
	status := ctx.GetHeader("status")
	var ctxt = context.Background()
	response, err := client.AdminViewPackages(ctxt, &pb.AdminView{
		Status: status,
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
		"message": fmt.Sprintf("packages fetched succesfully"),
		"data":    response,
	})
}

func ViewPackage(ctx *gin.Context, client pb.AdminClient) {
	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		fmt.Println("categoryID missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "pacakge id missing",
			"msg":    "error",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.AdminViewpackage(ctxt, &pb.AdminView{
		Id: int64(packageId),
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
		"message": fmt.Sprintf("packages fetched succesfully"),
		"data":    response,
	})
}

func PackageStatus(ctx *gin.Context, client pb.AdminClient) {
	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		fmt.Println("pacakgeID missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.AdminPacakgeStatus(ctxt, &pb.AdminView{
		Id: int64(packageId),
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
		"message": fmt.Sprintf("package updated succesfully"),
		"data":    response,
	})
}
