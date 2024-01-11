package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func CoordinatorLoginHandler(ctx *gin.Context, client cpb.CoordinatorClient, role string) {
	var login dto.Login

	if err := ctx.BindJSON(&login); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	err := validate.Struct(login)
	if err != nil {
		log.Printf("Validation error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Validation error",
		})
	}

	ctxt := context.Background()
	response, err := client.CoordinatorLoginRequest(ctxt, &cpb.Login{
		Email:    login.Email,
		Password: login.Password,
		Role:     role,
	})

	if err != nil {
		log.Printf("error logging in user %v err: %v", login.Email, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v logged in succesfully", login.Email),
		"data":    response,
	})
}

func CoordinatorSignupHandler(ctx *gin.Context, client cpb.CoordinatorClient) {
	var coordinator dto.User

	if err := ctx.BindJSON(&coordinator); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	err := validate.Struct(coordinator)
	if err != nil {
		log.Printf("Validation error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Validation error",
		})
	}

	ctxt := context.Background()
	response, err := client.CoordinatorSignupRequest(ctxt, &cpb.Signup{
		Name:     coordinator.Name,
		Email:    coordinator.Email,
		Phone:    int32(coordinator.Phone),
		Password: coordinator.Password,
	})

	if err != nil {
		log.Printf("error logging in user %v err: %v", coordinator.Email, err.Error())
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
