package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
)

func ViewBookings(ctx *gin.Context, client cpb.CoordinatorClient) {
	id := ctx.GetHeader("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("package id missing", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "package id missing",
		})
		return
	}
	var ctxt = context.Background()
	response, err := client.ViewHistory(ctxt, &cpb.View{
		Status: "true",
		Id:     int64(ID),
	})

	if err != nil {
		log.Printf("booking fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("booking fetched succesfully"),
		"data":    response,
	})
}

func ViewBooking(ctx *gin.Context, client cpb.CoordinatorClient) {
	id := ctx.GetHeader("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("package id missing", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	var ctxt = context.Background()
	response, err := client.ViewBooking(ctxt, &cpb.View{
		Id: int64(ID),
	})

	if err != nil {
		log.Printf("booking fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("booking fetched succesfully"),
		"data":    response,
	})
}

func ViewTraveller(ctx *gin.Context, client cpb.CoordinatorClient) {
	id := ctx.GetHeader("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("traveller id missing", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	var ctxt = context.Background()
	response, err := client.ViewTraveller(ctxt, &cpb.View{
		Id: int64(ID),
	})

	if err != nil {
		log.Printf("booking fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("booking fetched succesfully"),
		"data":    response,
	})
}