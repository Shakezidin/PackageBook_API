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

func AddActivity(ctx *gin.Context, client cpb.CoordinatorClient) {
	var activity dto.AddActivities

	destinationIdStr := ctx.GetHeader("id")
	destinationId, err := strconv.Atoi(destinationIdStr)
	if err != nil {
		fmt.Println("packageId missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	if err := ctx.BindJSON(&activity); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(activity)
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
	response, err := client.CoordinatorAddActivity(ctxt, &cpb.Activity{
		DestinationId: int64(destinationId),
		Activityname:  activity.ActivityName,
		Description:   activity.Description,
		Location:      activity.Location,
		ActivityType:  activity.ActivityType,
		Amount:        int64(activity.Price),
		Date:          activity.Date,
		Time:          activity.Time,
	})

	if err != nil {
		log.Printf("activity %s creattion error", activity.ActivityName, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v activity created succesfully", activity.ActivityName),
		"data":    response,
	})
}

func ViewActivity(ctx *gin.Context, client cpb.CoordinatorClient) {
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
	response, err := client.CoordinatorViewActivity(ctxt, &cpb.View{
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