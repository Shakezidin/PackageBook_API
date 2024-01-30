package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CoordinatorAddFoodMenu(ctx *gin.Context, client cpb.CoordinatorClient) {
	var foodmenu dto.FoodMenu

	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		fmt.Println("packageID missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "packageID missing",
		})
		return
	}

	if err := ctx.BindJSON(&foodmenu); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error binding JSON",
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(foodmenu)
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

	_, err = time.Parse("02-01-2006", foodmenu.Date)
	if err != nil {
		log.Printf("date fromat error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.CoordinatorAddFoodMenu(ctxt, &cpb.FoodMenu{
		PackageID: int64(packageId),
		Breakfast: foodmenu.Breakfast,
		Lunch:     foodmenu.Lunch,
		Dinner:    foodmenu.Dinner,
		Date:      foodmenu.Date,
	})

	if err != nil {
		log.Printf("food menu creattion error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("food menu created succesfully"),
		"data":    response,
	})
}

func CoordinatorViewFoodMenus(ctx *gin.Context, client cpb.CoordinatorClient) {
	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		fmt.Println("packageID missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "packageID missing",
		})
		return
	}
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	var ctxt = context.Background()
	response, err := client.CoordinatorViewFoodMenu(ctxt, &cpb.View{
		Id:   int64(packageId),
		Page: int64(pageInt),
	})

	if err != nil {
		log.Printf("food menu fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("food menu fetched succesfully"),
		"data":    response,
	})

}
