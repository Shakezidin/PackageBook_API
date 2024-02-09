package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	dto "github.com/Shakezidin/pkg/DTO"
	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AddCategory(ctx *gin.Context, client pb.AdminClient) {
	var category dto.AddCategory
	if err := ctx.BindJSON(&category); err != nil {
		errMsg := "error binding JSON"
		log.Println(errMsg, err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errMsg,
		})
		return
	}

	validate := validator.New()

	if err := validate.Struct(category); err != nil {
		errMsg := "validation error"
		log.Println(errMsg, err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errMsg,
		})
		return
	}

	ctxt := context.Background()
	response, err := client.AdminAddCategory(ctxt, &pb.AdminCategory{
		Category: category.Category,
	})

	if err != nil {
		errMsg := "error while adding category"
		log.Printf("%s %v: %v", errMsg, category.Category, err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errMsg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("%v added successfully", category.Category),
		"data":    response,
	})
}

func ViewCategories(ctx *gin.Context, client pb.AdminClient) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	ctxt := context.Background()
	response, err := client.AdminViewCategories(ctxt, &pb.AdminView{
		Page: int64(page),
	})

	if err != nil {
		errMsg := "error while fetching categories"
		log.Printf("%s: %v", errMsg, err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errMsg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "categories fetched successfully",
		"data":    response,
	})
}
