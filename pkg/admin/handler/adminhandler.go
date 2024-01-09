package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	dto "github.com/Shakezidin/pkg/DTO"
	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AdminLoginHandler(ctx *gin.Context, client pb.AdminClient, role string) {
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
	}

	ctxt := context.Background()
	response, err := client.AdminLoginRequest(ctxt, &pb.Login{
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
