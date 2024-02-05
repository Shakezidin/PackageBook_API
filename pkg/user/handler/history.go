package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Shakezidin/middleware"
	pb "github.com/Shakezidin/pkg/user/pb"
	"github.com/gin-gonic/gin"
)

func ViewHistory(ctx *gin.Context, client pb.UserClient) {
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	var ctxt = context.Background()
	_, id, err := middleware.ValidateToken(ctx, "user")
	userId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("token validation error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	response, err := client.UserViewHistory(ctxt, &pb.UserView{
		Page: int64(pageInt),
		Id:   int64(userId),
	})

	if err != nil {
		log.Printf("history fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("history  fetched succesfully"),
		"data":    response,
	})
}

func ViewBooking(ctx *gin.Context, client pb.UserClient) {
	var ctxt = context.Background()
	id := ctx.GetHeader("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("token validation error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	response, err := client.UserViewBooking(ctxt, &pb.UserView{
		Id: int64(Id),
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
		"message": fmt.Sprintf("booking  fetched succesfully"),
		"data":    response,
	})
}

func PackageCancel(ctx *gin.Context, client pb.UserClient) {
	id := ctx.GetHeader("id")
	startdate := ctx.GetHeader("startdate")
	Id, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Error booking id missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "booking id missing",
		})
		return
	}

	strtdate, err := time.Parse("02-01-2006", startdate)
	if err != nil {
		log.Printf("Error while date passing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Error while date passing",
		})
		return
	}

	duration := strtdate.Sub(time.Now())
	if duration <= 24*time.Hour {
		log.Printf("cannot cancel the package")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "cancellation time out",
		})
		return
	}

	_, userid, err := middleware.ValidateToken(ctx, "user")
	if err != nil {
		log.Printf("token validation error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	userId, _ := strconv.Atoi(userid)
	var ctxt = context.Background()
	response, err := client.UserCancelBooking(ctxt, &pb.UserView{
		Id:     int64(Id),
		UserId: int64(userId),
	})

	if err != nil {
		log.Printf("Error while cancelling package", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("Package cancelled successfully"),
		"data":    response,
	})
}
