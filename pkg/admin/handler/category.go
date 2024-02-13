package handler

import (
	"context"
	"net/http"
	"strconv"

	dto "github.com/Shakezidin/pkg/DTO"
	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AddCategory handles the HTTP request to add a new category.
func AddCategory(ctx *gin.Context, client pb.AdminClient) {
	var category dto.AddCategory
	if err := ctx.BindJSON(&category); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Invalid request body",
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(category); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "Validation failed",
		})
		return
	}

	ctxt := context.Background()
	response, err := client.AdminAddCategory(ctxt, &pb.AdminCategory{
		Category: category.Category,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Category added successfully",
		"Data":    response,
	})
}

// ViewCategories handles the HTTP request to view categories.
func ViewCategories(ctx *gin.Context, client pb.AdminClient) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	ctxt := context.Background()
	response, err := client.AdminViewCategories(ctxt, &pb.AdminView{
		Page: int64(page),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Categories fetched successfully",
		"Data":    response,
	})
}
