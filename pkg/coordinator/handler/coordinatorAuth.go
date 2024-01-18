package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	utility "github.com/Shakezidin/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// var validate = validator.New()

func CoordinatorSignupHandler(ctx *gin.Context, client cpb.CoordinatorClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var coordinator dto.User

	if err := ctx.BindJSON(&coordinator); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	//? Validating struct
	validate := validator.New()
	validate.RegisterValidation("emailcst", utility.EmailValidation)
	validate.RegisterValidation("phone", utility.PhoneNumberValidation)
	validate.RegisterValidation("alphaspace", utility.AlphaSpace)
	err := validate.Struct(coordinator)
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

	response, err := client.CoordinatorSignupRequest(cont, &cpb.Signup{
		Name:     coordinator.Name,
		Email:    coordinator.Email,
		Phone:    coordinator.Phone,
		Password: coordinator.Password,
		Role:     "coordinator",
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

func VerifySignup(ctx *gin.Context, client cpb.CoordinatorClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var req dto.Otp
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response, err := client.CoordinatorSignupVerifyRequest(cont, &cpb.Verify{
		OTP:   int32(req.OTP),
		Email: req.Email,
	})

	if err != nil {
		log.Printf("error registering user err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusAccepted,
		"message": "OTP verified, user Creation successful.",
		"data":    response,
	})
}

func CoordinatorLoginHandler(ctx *gin.Context, client cpb.CoordinatorClient, role string) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var login dto.Login

	if err := ctx.BindJSON(&login); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(login)
	if err != nil {
		log.Printf("Validation error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Validation error",
		})
		return
	}

	response, err := client.CoordinatorLoginRequest(cont, &cpb.Login{
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
		"message": fmt.Sprintf("%v logged  in succesfully", login.Email),
		"data":    response,
	})
}
