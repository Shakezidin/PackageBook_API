package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Shakezidin/middleware"
	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ForgetPassword(ctx *gin.Context, client cpb.CoordinatorClient) {
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

	response, err := client.CoordinatorForgetPassword(cont, &cpb.ForgetPassword{
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

func ForgetPasswordVerify(ctx *gin.Context, client cpb.CoordinatorClient) {
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

	response, err := client.CoordinatorForgetPasswordVerify(cont, &cpb.ForgetPasswordVerify{
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

func NewPassword(ctx *gin.Context, client cpb.CoordinatorClient) {
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

	_, userId, err := middleware.ValidateToken(ctx, "coordinator")
	if err != nil {
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

	response, err := client.CoordinatorNewPassword(cont, &cpb.Newpassword{
		Newpassword: newpassword.NewPassword,
		Id:          userId,
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

func ViewDashBord(ctx *gin.Context, client cpb.CoordinatorClient) {
	_, id, err := middleware.ValidateToken(ctx, "coordinator")
	if err != nil {
		log.Printf("Token validation error, err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	Id, _ := strconv.Atoi(id)
	response, err := client.ViewDashBord(ctx, &cpb.View{
		Id: int64(Id),
	})

	if err != nil {
		log.Printf("error while fetching dashbord err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status": http.StatusOK,
		"data":   response,
	})
}
