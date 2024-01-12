package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Shakezidin/middleware"
	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AddPackage(ctx *gin.Context, client cpb.CoordinatorClient) {
	var pkg dto.Addpackage

	categoryIdStr := ctx.GetHeader("id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		fmt.Println("categoryId missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	email, err := middleware.ValidateToken(ctx, "coordinator")
	if err != nil {
		log.Printf("token validation error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	if err := ctx.BindJSON(&pkg); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(pkg)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		for _, e := range err.(validator.ValidationErrors) {
			log.Printf("struct validation errors %v, %v", e.Field(), e.Tag())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("error in field %v, error: %v", e.Field(), e.Tag()),
			})
		}
		return
	}

	var ctxt = context.Background()
	response, err := client.CoordinatorAddPackage(ctxt, &cpb.AddPackage{
		Packagename:      pkg.Name,
		Coorinatoremail:  email,
		Startdate:        pkg.StartDateTime,
		Startlocation:    pkg.StartLocation,
		Enddate:          pkg.EndDateTime,
		Endlocation:      pkg.EndLocation,
		Price:            int64(pkg.Price),
		Image:            pkg.Image,
		DestinationCount: int64(pkg.DestinationCount),
		Destination:      pkg.Destination,
		CategoryId:       int64(categoryId),
		Description:      pkg.Description,
		MaxCapacity:      pkg.MaxCapacity,
	})

	if err != nil {
		log.Printf("pakcage %s creattion error", pkg.Name, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v package created succesfully", pkg.Name),
		"data":    response,
	})

}

func AddDestination(ctx *gin.Context, client cpb.CoordinatorClient) {
	var destination dto.AddDestination

	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		fmt.Println("packageId missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	if err := ctx.BindJSON(&destination); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(destination)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		for _, e := range err.(validator.ValidationErrors) {
			log.Printf("struct validation errors %v, %v", e.Field(), e.Tag())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("error in field %v, error: %v", e.Field(), e.Tag()),
			})
		}
		return
	}

	var ctxt = context.Background()
	response, err := client.CoordinatorAddDestination(ctxt, &cpb.AddDestination{
		DestinationName: destination.DestinationName,
		Description:     destination.Description,
		PackageId:       int64(packageId),
		Minprice:        int64(destination.MinPrice),
		Image:           destination.Image,
		MaxCapacity:     int64(destination.MaxCapacity),
	})

	if err != nil {
		log.Printf("destination %s creattion error", destination.DestinationName, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v destination created succesfully", destination.DestinationName),
		"data":    response,
	})
}
