package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	utility "github.com/Shakezidin/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CoordinatorSignupHandler handles the coordinator signup request.
func CoordinatorSignupHandler(ctx *gin.Context, client cpb.CoordinatorClient) {
	// Set timeout for the context
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var coordinator dto.User

	// Bind request body to coordinator struct
	if err := ctx.BindJSON(&coordinator); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	// Validate the coordinator struct
	validate := validator.New()
	validate.RegisterValidation("email", utility.EmailValidation)
	validate.RegisterValidation("phone", utility.PhoneNumberValidation)
	validate.RegisterValidation("alphaspace", utility.AlphaSpace)
	err := validate.Struct(coordinator)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
		})
		for _, e := range err.(validator.ValidationErrors) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Error": fmt.Sprintf("Error in field %v, error: %v", e.Field(), e.Tag()),
			})
		}
		return
	}

	// Make gRPC call to signup
	response, err := client.CoordinatorSignupRequest(cont, &cpb.Signup{
		Name:     coordinator.Name,
		Email:    coordinator.Email,
		Phone:    coordinator.Phone,
		Password: coordinator.Password,
		Role:     "coordinator",
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status": http.StatusOK,
		"Data":   response,
	})
}

// VerifySignup verifies the coordinator signup OTP.
func VerifySignup(ctx *gin.Context, client cpb.CoordinatorClient) {
	// Set timeout for the context
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var req dto.Otp

	// Bind request body to OTP struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	// Make gRPC call to verify OTP
	response, err := client.CoordinatorSignupVerifyRequest(cont, &cpb.Verify{
		OTP:   int32(req.OTP),
		Email: req.Email,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Status":  http.StatusCreated,
		"Message": "OTP verified, coordinator creation successful.",
		"Data":    response,
	})
}

// CoordinatorLoginHandler handles the coordinator login request.
func CoordinatorLoginHandler(ctx *gin.Context, client cpb.CoordinatorClient, role string) {
	// Set timeout for the context
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var login dto.Login

	// Bind request body to login struct
	if err := ctx.BindJSON(&login); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	// Validate the login struct
	validate := validator.New()
	err := validate.Struct(login)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Validation error",
		})
		return
	}

	// Make gRPC call to login
	response, err := client.CoordinatorLoginRequest(cont, &cpb.Login{
		Email:    login.Email,
		Password: login.Password,
		Role:     role,
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
		"Message": fmt.Sprintf("%v logged  in successfully", login.Email),
		"Data":    response,
	})
}
