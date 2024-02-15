package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AddActivity(ctx *gin.Context, client cpb.CoordinatorClient) {
	var activity dto.AddActivities

	destinationIdStr := ctx.GetHeader("id")
	destinationId, err := strconv.Atoi(destinationIdStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	if err := ctx.BindJSON(&activity); err != nil {
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
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("error in field %v, error: %v", e.Field(), e.Tag()),
			})
		}
		return
	}

	var ctxt = context.Background()
	response, err := client.CoordinatorAddActivity(ctxt, &cpb.Activity{
		Activityname:  activity.ActivityName,
		Description:   activity.Description,
		Location:      activity.Location,
		ActivityType:  activity.ActivityType,
		Amount:        int64(activity.Price),
		Date:          activity.Date,
		Time:          activity.Time,
		DestinationId: int64(destinationId),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status": http.StatusAccepted,
		"data":   response,
	})
}

func ViewActivity(ctx *gin.Context, client cpb.CoordinatorClient) {
	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)

	if err != nil {
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": "activity fetched succesfully",
		"data":    response,
	})
}
