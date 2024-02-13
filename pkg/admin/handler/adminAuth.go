package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Shakezidin/middleware"
	dto "github.com/Shakezidin/pkg/DTO"
	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AdminLoginHandler handles the HTTP request for admin login.
func AdminLoginHandler(ctx *gin.Context, client pb.AdminClient, role string) {
	// Bind the JSON request body to the Login struct.
	var login dto.Login
	if err := ctx.BindJSON(&login); err != nil {
		errMsg := "Error binding JSON"
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errMsg,
		})
		return
	}

	// Validate the request body using the validator.
	validate := validator.New()
	if err := validate.Struct(login); err != nil {
		errMsg := "Validation error"
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  errMsg,
		})
		return
	}

	// Create a context for the gRPC request.
	ctxt := context.Background()

	// Call the gRPC service to log in the admin user.
	response, err := client.AdminLoginRequest(ctxt, &pb.AdminLogin{
		Email:    login.Email,
		Password: login.Password,
		Role:     role,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Respond with the successful login message and data.
	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": fmt.Sprintf("%s logged in successfully", login.Email),
		"Data":    response,
	})
}

// ViewDashboard handles the HTTP request to view the admin dashboard.
func ViewDashboard(ctx *gin.Context, client pb.AdminClient) {
	// Validate the token and extract the user's email.
	email, _, err := middleware.ValidateToken(ctx, "admin")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Status": http.StatusUnauthorized,
			"Error":  err.Error(),
		})
		return
	}

	// Call the gRPC service to fetch the admin dashboard.
	response, err := client.AdminViewDashBord(ctx, &pb.AdminView{Status: email})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Respond with the dashboard data.
	ctx.JSON(http.StatusOK, gin.H{
		"Status": http.StatusOK,
		"Data":   response,
	})
}
