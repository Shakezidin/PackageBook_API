package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Shakezidin/middleware"
	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ForgetPassword initiates the process of resetting a user's password.
func ForgetPassword(ctx *gin.Context, client cpb.CoordinatorClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var frgtpswrd dto.ForgetPassword
	if err := ctx.BindJSON(&frgtpswrd); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Error binding JSON",
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(frgtpswrd)
	if err != nil {
		handleValidationError(ctx, err)
		return
	}

	response, err := client.CoordinatorForgetPassword(cont, &cpb.ForgetPassword{
		Phone: frgtpswrd.Phone,
	})

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

// ForgetPasswordVerify verifies the OTP sent to the user for password reset.
func ForgetPasswordVerify(ctx *gin.Context, client cpb.CoordinatorClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var otp dto.PswrdOtp
	if err := ctx.BindJSON(&otp); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Error binding JSON",
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(otp)
	if err != nil {
		handleValidationError(ctx, err)
		return
	}

	response, err := client.CoordinatorForgetPasswordVerify(cont, &cpb.ForgetPasswordVerify{
		OTP:   otp.OTP,
		Phone: otp.Phone,
	})

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

// NewPassword sets the new password after successful verification of OTP.
func NewPassword(ctx *gin.Context, client cpb.CoordinatorClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var newpassword dto.NewChange
	if err := ctx.BindJSON(&newpassword); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Error binding JSON",
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(newpassword)
	if err != nil {
		handleValidationError(ctx, err)
		return
	}

	_, userId, err := middleware.ValidateToken(ctx, "coordinator")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Token validation error",
		})
		return
	}

	if newpassword.ConfirmPassword != newpassword.NewPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Password mismatch",
		})
		return
	}

	response, err := client.CoordinatorNewPassword(cont, &cpb.Newpassword{
		New_Password: newpassword.NewPassword,
		ID:          userId,
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

// ViewDashboard retrieves the dashboard data for a coordinator.
func ViewDashboard(ctx *gin.Context, client cpb.CoordinatorClient) {
	_, id, err := middleware.ValidateToken(ctx, "coordinator")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	Id, _ := strconv.Atoi(id)
	response, err := client.ViewDashboard(ctx, &cpb.View{
		ID: int64(Id),
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

// handleValidationError handles validation errors.
func handleValidationError(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"Status": http.StatusBadRequest,
	})
	for _, e := range err.(validator.ValidationErrors) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": fmt.Sprintf("Error in field %v, error: %v", e.Field(), e.Tag()),
		})
	}
}
