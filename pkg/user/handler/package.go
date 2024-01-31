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
		fmt.Println("package ID missing")
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

func ViewCatagories(ctx *gin.Context, client pb.UserClient) {
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	var ctxt = context.Background()
	response, err := client.UserViewCatagories(ctxt, &pb.UserView{
		Page: int64(pageInt),
	})

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

func ViewPackages(ctx *gin.Context, client pb.UserClient) {
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	var ctxt = context.Background()
	response, err := client.UserViewPackages(ctxt, &pb.UserView{
		Status: true,
		Page:   int64(pageInt),
	})

	if err != nil {
		log.Printf("packages fetching  error", err.Error())
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

func ViewFoodMenus(ctx *gin.Context, client pb.UserClient) {
	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		fmt.Println("packageID missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "packageID missing",
		})
		return
	}
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	var ctxt = context.Background()
	response, err := client.UserViewFoodMenu(ctxt, &pb.UserView{
		Id:   int64(packageId),
		Page: int64(pageInt),
	})

	if err != nil {
		log.Printf("food menu fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("food menu fetched succesfully"),
		"data":    response,
	})
}

type Filter struct {
	StartTime  string  `json:"starttime"`
	MinPrice   float64 `json:"minprice" validate:"min=5"`
	MaxPrice   float64 `json:"maxprice" validate:"max=10000"`
	Orderby    string  `json:"orderby"`
	CategoryId string  `json:"categoryid"`
}

func PackageFilter(ctx *gin.Context, client pb.UserClient) {
	var filter Filter

	if err := ctx.BindJSON(&filter); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	categoryId, _ := strconv.Atoi(filter.CategoryId)
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	var ctxt = context.Background()
	response, err := client.UserFilterPackage(ctxt, &pb.UserFilter{
		Status:       true,
		Page:         int64(pageInt),
		MinPrice:     int64(filter.MinPrice),
		MaxPrice:     int64(filter.MaxPrice),
		OrderBy:      filter.Orderby,
		CategoryId:   int64(categoryId),
		Departurtime: filter.StartTime,
	})

	if err != nil {
		log.Printf("packages fetching  error", err.Error())
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
