package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Shakezidin/middleware"
	dto "github.com/Shakezidin/pkg/DTO"
	pb "github.com/Shakezidin/pkg/user/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ForgetPassword(ctx *gin.Context, client pb.UserClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var frgtpswrd dto.ForgetPassword

	if err := ctx.BindJSON(&frgtpswrd); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(frgtpswrd)
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

	response, err := client.ForgetPassword(cont, &pb.ForgetPassword{
		Phone: frgtpswrd.Phone,
	})

	if err != nil {
		log.Printf("error senting otp in user %v err: %v", frgtpswrd.Phone, err.Error())
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

func ForgetPasswordVerify(ctx *gin.Context, client pb.UserClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var otp dto.PswrdOtp

	if err := ctx.BindJSON(&otp); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(otp)
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

	response, err := client.ForgetPasswordVerify(cont, &pb.ForgetPasswordVerify{
		Otp:   otp.OTP,
		Phone: otp.Phone,
	})

	if err != nil {
		log.Printf("error senting otp in phone %v err: %v", otp.Phone, err.Error())
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

func NewPassword(ctx *gin.Context, client pb.UserClient) {
	timeout := time.Second * 1000
	cont, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var newpassword dto.NewChange

	if err := ctx.BindJSON(&newpassword); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(newpassword)
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

	_,userId,err:=middleware.ValidateToken(ctx,"user")
	if err!=nil{
		fmt.Println("tocken validation error")
		ctx.JSON(200, gin.H{
			"status": http.StatusAccepted,
			"data":   "tocken validation error",
		})
		return
	}

	if newpassword.ConfirmPassword != newpassword.NewPassword {
		fmt.Println("password missmach")
		ctx.JSON(200, gin.H{
			"status": http.StatusAccepted,
			"data":   "massword missmatch",
		})
		return
	}

	response, err := client.NewPassword(cont, &pb.Newpassword{
		Newpassword: newpassword.NewPassword,
		Id: userId,
	})

	if err != nil {
		log.Printf("error setting new password err: %v", err.Error())
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
