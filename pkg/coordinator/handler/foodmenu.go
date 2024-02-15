package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CoordinatorAddFoodMenu handles the addition of a new food menu for a package.
func CoordinatorAddFoodMenu(ctx *gin.Context, client cpb.CoordinatorClient) {
	var foodmenu dto.FoodMenu

	packageIDStr := ctx.GetHeader("id")
	packageID, err := strconv.Atoi(packageIDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Package ID missing",
		})
		return
	}

	if err := ctx.BindJSON(&foodmenu); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Error binding JSON",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(foodmenu); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Validation error",
		})
		for _, e := range err.(validator.ValidationErrors) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Error": fmt.Sprintf("Error in field %v, error: %v", e.Field(), e.Tag()),
			})
		}
		return
	}

	_, err = time.Parse("02-01-2006", foodmenu.Date)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Date format error",
		})
		return
	}

	ctxt := context.Background()
	response, err := client.CoordinatorAddFoodMenu(ctxt, &cpb.FoodMenu{
		PackageID: int64(packageID),
		Breakfast: foodmenu.Breakfast,
		Lunch:     foodmenu.Lunch,
		Dinner:    foodmenu.Dinner,
		Date:      foodmenu.Date,
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
		"Message": "Food menu created successfully",
		"Data":    response,
	})
}

// CoordinatorViewFoodMenus fetches food menus for a specific package.
func CoordinatorViewFoodMenus(ctx *gin.Context, client cpb.CoordinatorClient) {
	packageIDStr := ctx.GetHeader("id")
	packageID, err := strconv.Atoi(packageIDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "PackageID missing",
		})
		return
	}

	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)

	ctxt := context.Background()
	response, err := client.CoordinatorViewFoodMenu(ctxt, &cpb.View{
		Id:   int64(packageID),
		Page: int64(pageInt),
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
		"Message": "Food menu fetched successfully",
		"Data":    response,
	})
}
