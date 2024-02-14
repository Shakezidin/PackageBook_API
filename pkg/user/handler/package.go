package handler

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

// ViewPackage retrieves details of a specific package.
func ViewPackage(ctx *gin.Context, client pb.UserClient) {
	packageIDStr := ctx.GetHeader("id")
	packageID, err := strconv.Atoi(packageIDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Package ID missing",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.UserViewPackage(ctxt, &pb.UserView{
		Id: int64(packageID),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Package fetched successfully",
		"Data":    response,
	})
}

// ViewCategories retrieves a list of categories.
func ViewCategories(ctx *gin.Context, client pb.UserClient) {
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	var ctxt = context.Background()
	response, err := client.UserViewCategories(ctxt, &pb.UserView{
		Page: int64(pageInt),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Categories fetched successfully",
		"Data":    response,
	})
}

// ViewPackages retrieves a list of packages.
func ViewPackages(ctx *gin.Context, client pb.UserClient) {
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	var ctxt = context.Background()
	response, err := client.UserViewPackages(ctxt, &pb.UserView{
		Status: "true",
		Page:   int64(pageInt),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Packages fetched successfully",
		"Data":    response,
	})
}

// ViewFoodMenus retrieves food menus for a specific package.
func ViewFoodMenus(ctx *gin.Context, client pb.UserClient) {
	packageIDStr := ctx.GetHeader("id")
	packageID, err := strconv.Atoi(packageIDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"Error":  "Package ID missing",
		})
		return
	}
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	var ctxt = context.Background()
	response, err := client.UserViewFoodMenu(ctxt, &pb.UserView{
		Id:   int64(packageID),
		Page: int64(pageInt),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Food menu fetched successfully",
		"Data":    response,
	})
}

type Filter struct {
	StartTime  string  `json:"starttime"`
	MinPrice   float64 `json:"minprice" validate:"min=5"`
	MaxPrice   float64 `json:"maxprice" validate:"max=10000"`
	Orderby    string  `json:"orderby"`
	CategoryID string  `json:"categoryid"`
}

// PackageFilter filters packages based on various criteria.
func PackageFilter(ctx *gin.Context, client pb.UserClient) {
	var filter Filter

	if err := ctx.BindJSON(&filter); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Binding error",
		})
		return
	}

	categoryID, _ := strconv.Atoi(filter.CategoryID)
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	var ctxt = context.Background()
	response, err := client.UserFilterPackage(ctxt, &pb.UserFilter{
		Status:       true,
		Page:         int64(pageInt),
		MinPrice:     int64(filter.MinPrice),
		MaxPrice:     int64(filter.MaxPrice),
		OrderBy:      filter.Orderby,
		CategoryId:   int64(categoryID),
		Departurtime: filter.StartTime,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Packages fetched successfully",
		"Data":    response,
	})
}
