package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

// ViewPackages handles the HTTP request to view packages.
func ViewPackages(ctx *gin.Context, client pb.AdminClient) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	status := ctx.GetHeader("status")

	// Create a context
	ctxt := context.Background()

	// Call the gRPC service to fetch packages
	response, err := client.AdminViewPackages(ctxt, &pb.AdminView{
		Status: status,
		Page:   int64(page),
	})

	// Handle errors
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Respond with fetched packages
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Packages fetched successfully",
		"Data":    response,
	})
}

// ViewPackage handles the HTTP request to view a specific package.
func ViewPackage(ctx *gin.Context, client pb.AdminClient) {
	// Retrieve package ID from header
	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		log.Printf("Package ID missing or invalid: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Package ID missing or invalid",
		})
		return
	}

	// Create a context
	ctxt := context.Background()

	// Call the gRPC service to fetch the package
	response, err := client.AdminViewpackage(ctxt, &pb.AdminView{
		Id: int64(packageId),
	})

	// Handle errors
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Respond with fetched package
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Package fetched successfully",
		"Data":    response,
	})
}

// PackageStatus handles the HTTP request to update the status of a package.
func PackageStatus(ctx *gin.Context, client pb.AdminClient) {
	// Retrieve package ID from header
	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Package ID missing or invalid",
		})
		return
	}

	// Create a context
	ctxt := context.Background()

	// Call the gRPC service to update the package status
	response, err := client.AdminPacakgeStatus(ctxt, &pb.AdminView{
		Id: int64(packageId),
	})

	// Handle errors
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Respond with updated package status
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Package status updated successfully",
		"Data":    response,
	})
}
