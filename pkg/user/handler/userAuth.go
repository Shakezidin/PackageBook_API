package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	dto "github.com/Shakezidin/pkg/DTO"
	pb "github.com/Shakezidin/pkg/user/userpb"
	utility "github.com/Shakezidin/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// UserSignupHandler handles the user signup request.
func UserSignupHandler(ctx *gin.Context, client pb.UserClient, role string) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var user dto.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	// Validate struct
	validate := validator.New()
	validate.RegisterValidation("emailcst", utility.EmailValidation)
	validate.RegisterValidation("phone", utility.PhoneNumberValidation)
	validate.RegisterValidation("alphaspace", utility.AlphaSpace)
	err := validate.Struct(user)
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

	response, err := client.UserSignupRequest(cont, &pb.UserSignup{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
		Role:     role,
	})
	fmt.Println(err)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"Status": http.StatusAccepted,
		"Data":   response,
	})
}

// VerifySignup verifies the user signup request.
func VerifySignup(ctx *gin.Context, client pb.UserClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var req dto.Otp
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	response, err := client.UserSignupVerifyRequest(cont, &pb.UserVerify{
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
		"Message": "OTP verified, user creation successful.",
		"Data":    response,
	})
}

// UserLoginHandler handles the user login request.
func UserLoginHandler(ctx *gin.Context, client pb.UserClient, role string) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var login dto.Login

	if err := ctx.BindJSON(&login); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(login)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Validation error",
		})
		return
	}

	response, err := client.UserLoginRequest(cont, &pb.UserLogin{
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
		"Message": fmt.Sprintf("%v logged in successfully", login.Email),
		"Data":    response,
	})
}
