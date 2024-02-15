package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Shakezidin/middleware"
	dto "github.com/Shakezidin/pkg/DTO"
	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ForgetPassword initiates the process to reset the user's password.
func ForgetPassword(ctx *gin.Context, client pb.UserClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var frgtpswrd dto.ForgetPassword
	if err := ctx.BindJSON(&frgtpswrd); err != nil {
		handleError(ctx, err, "Error binding JSON")
		return
	}

	validate := validator.New()
	if err := validate.Struct(frgtpswrd); err != nil {
		handleValidationErrors(ctx, err)
		return
	}

	response, err := client.UserForgetPassword(cont, &pb.UserforgetPassword{
		Phone: frgtpswrd.Phone,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"Status": http.StatusAccepted,
		"message":"password change initiated, check message for OTP",
		"Data":   response,
	})
}

// ForgetPasswordVerify verifies the OTP for resetting the user's password.
func ForgetPasswordVerify(ctx *gin.Context, client pb.UserClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var otp dto.PswrdOtp
	if err := ctx.BindJSON(&otp); err != nil {
		handleError(ctx, err, "Error binding JSON")
		return
	}

	validate := validator.New()
	if err := validate.Struct(otp); err != nil {
		handleValidationErrors(ctx, err)
		return
	}

	response, err := client.UserForgetPasswordVerify(cont, &pb.UserforgetPasswordVerify{
		Otp:   otp.OTP,
		Phone: otp.Phone,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"Status": http.StatusAccepted,
		"Data":   response,
	})
}

// NewPassword sets the new password for the user.
func NewPassword(ctx *gin.Context, client pb.UserClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var newPassword dto.NewChange
	if err := ctx.BindJSON(&newPassword); err != nil {
		handleError(ctx, err, "Error binding JSON")
		return
	}

	validate := validator.New()
	if err := validate.Struct(newPassword); err != nil {
		handleValidationErrors(ctx, err)
		return
	}

	_, userID, err := middleware.ValidateToken(ctx, "user")
	if err != nil {
		handleError(ctx, err, "Token validation error")
		return
	}

	if newPassword.ConfirmPassword != newPassword.NewPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Password mismatch",
		})
		return
	}

	response, err := client.UserNewPassword(cont, &pb.Usernewpassword{
		Newpassword: newPassword.NewPassword,
		Id:          userID,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status": http.StatusOK,
		"Data":   response,
	})
}

// UpdateProfile updates the user's profile details.
func UpdateProfile(ctx *gin.Context, client pb.UserClient) {
	var updateUser dto.User
	if err := ctx.BindJSON(&updateUser); err != nil {
		handleError(ctx, err, "Error binding JSON")
		return
	}

	_, userID, err := middleware.ValidateToken(ctx, "user")
	if err != nil {
		handleError(ctx, err, "Token validation error")
		return
	}

	userIDInt, _ := strconv.Atoi(userID)
	response, err := client.UserProfileUpdate(context.Background(), &pb.UserSignup{
		Id:    int64(userIDInt),
		Name:  updateUser.Name,
		Email: updateUser.Email,
		Phone: updateUser.Phone,
		Role:  "user",
	})
	if err != nil {
		handleError(ctx, err, "Error updating profile")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status": http.StatusOK,
		"Data":   response,
	})
}

// handleError logs the error and responds with a JSON error message.
func handleError(ctx *gin.Context, err error, msg string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"Status": http.StatusBadRequest,
		"Error":  err.Error(),
	})
}

// handleValidationErrors responds with JSON error messages for validation errors.
func handleValidationErrors(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"Status": http.StatusBadRequest,
		"Error":  "Validation errors",
	})
}
