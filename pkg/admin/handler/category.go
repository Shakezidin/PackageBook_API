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

func AddCategory(ctx *gin.Context, client pb.AdminClient) {
	var category dto.AddCategory
	if err := ctx.BindJSON(&category); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	validate := validator.New()

	err := validate.Struct(category)
	if err != nil {
		log.Printf("Validation error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Validation error",
		})
		return
	}

	ctxt := context.Background()
	response, err := client.AdminAddCategory(ctxt, &pb.AdminCategory{
		Category: category.Category,
	})

	if err != nil {
		log.Printf("error while adding category %v err: %v", category.Category, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v added success", category.Category),
		"data":    response,
	})

}

func ViewCatagories(ctx *gin.Context, client pb.AdminClient) {

	ctxt := context.Background()
	response, err := client.AdminViewCategories(ctxt, &pb.AdminView{})

	if err != nil {
		log.Printf("error while viewing category , err: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("catagories fetched success"),
		"data":    response,
	})

}
