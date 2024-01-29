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
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	status := ctx.GetHeader("status")
	var ctxt = context.Background()
	response, err := client.AdminViewPackages(ctxt, &pb.AdminView{
		Status: status,
		Page:   int64(page),
	})

	if err != nil {
		log.Printf("error while fetching packages", err.Error())
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
		fmt.Println("packageID missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "package id missing",
			"msg":    "error",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.AdminViewpackage(ctxt, &pb.AdminView{
		Id: int64(packageId),
	})

	if err != nil {
		log.Printf("error while fetching package", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusOK,
		"message": "package fetched succesfully",
		"data":    response,
	})
}

func PackageStatus(ctx *gin.Context, client pb.AdminClient) {
	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		fmt.Println("packageID missing")
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
		log.Printf("error while updating package status", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("package status updated succesfully"),
		"data":    response,
	})
}
