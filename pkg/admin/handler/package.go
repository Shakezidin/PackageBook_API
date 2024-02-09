package handler

import (
	"context"
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
		log.Printf("error while fetching packages: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "failed to fetch packages",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "packages fetched successfully",
		"data":    response,
	})
}

func ViewPackage(ctx *gin.Context, client pb.AdminClient) {
	packageIdStr := ctx.GetHeader("id") // Retrieve package ID from path parameter
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		log.Printf("package ID missing or invalid: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "package ID missing or invalid",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.AdminViewpackage(ctxt, &pb.AdminView{
		Id: int64(packageId),
	})

	if err != nil {
		log.Printf("error while fetching package: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "failed to fetch package",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "package fetched successfully",
		"data":    response,
	})
}

func PackageStatus(ctx *gin.Context, client pb.AdminClient) {
	packageIdStr := ctx.GetHeader("id") // Retrieve package ID from path parameter
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		log.Printf("package ID missing or invalid: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "package ID missing or invalid",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.AdminPacakgeStatus(ctxt, &pb.AdminView{
		Id: int64(packageId),
	})

	if err != nil {
		log.Printf("error while updating package status: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "failed to update package status",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "package status updated successfully",
		"data":    response,
	})
}
