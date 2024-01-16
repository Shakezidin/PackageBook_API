package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	dto "github.com/Shakezidin/pkg/DTO"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
)

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
	response, err := client.CoordinatorAddDestination(ctxt, &cpb.Destination{
		DestinationName: destination.DestinationName,
		Description:     destination.Description,
		PackageID:       int64(packageId),
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

func ViewDestination(ctx *gin.Context, client cpb.CoordinatorClient) {
	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		fmt.Println("destination missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.CoordinatorViewDestination(ctxt, &cpb.View{
		Id: int64(packageId),
	})

	if err != nil {
		log.Printf("destination fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("destination fetched succesfully"),
		"data":    response,
	})
}
