package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Shakezidin/middleware"
	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AddPackage handles the addition of a new package.
func AddPackage(ctx *gin.Context, client cpb.CoordinatorClient) {
	var pkg dto.Addpackage

	destination := ctx.GetHeader("destination")
	startLocation := ctx.GetHeader("startlocation")
	categoryIDStr := ctx.GetHeader("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Category ID missing",
		})
		return
	}

	_, ID, err := middleware.ValidateToken(ctx, "coordinator")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	if err := ctx.BindJSON(&pkg); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Error while binding JSON",
		})
		return
	}
	pkg.Destination = destination
	pkg.StartLocation = startLocation

	validate := validator.New()
	err = validate.Struct(pkg)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
		})
		for _, e := range err.(validator.ValidationErrors) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Error": fmt.Sprintf("Error in field %v, error: %v", e.Field(), e.Tag()),
			})
		}
		return
	}

	id, _ := strconv.Atoi(ID)

	startDate, err := time.Parse("02-01-2006", pkg.StartDate)
	_, err = time.Parse("03:04 PM", pkg.StartTime)
	endDate, err := time.Parse("02-01-2006", pkg.EndDate)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"Error":  "Date format error",
		})
		return
	}

	if !endDate.Add(24 * time.Hour).After(startDate) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Error in start date and end date",
		})
		return
	}

	ctxt := context.Background()
	response, err := client.CoordinatorAddPackage(ctxt, &cpb.Package{
		Coorinator_ID:     int64(id),
		Package_Name:      pkg.Name,
		Start_Date:        pkg.StartDate,
		Start_Time:        pkg.StartTime,
		Start_Location:    pkg.StartLocation,
		End_Date:          pkg.EndDate,
		Price:            int64(pkg.Price),
		Image:            pkg.Image,
		Destination_Count: int64(pkg.DestinationCount),
		Destination:      pkg.Destination,
		Category_ID:       int64(categoryID),
		Description:      pkg.Description,
		Max_Capacity:      pkg.MaxCapacity,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Status": http.StatusCreated,
		"Data":   response,
	})
}

// ViewPackage fetches a package by ID.
func ViewPackage(ctx *gin.Context, client cpb.CoordinatorClient) {
	packageIDStr := ctx.GetHeader("id")
	packageID, err := strconv.Atoi(packageIDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Package ID missing",
		})
		return
	}

	ctxt := context.Background()
	response, err := client.CoordinatorViewPackage(ctxt, &cpb.View{
		ID: int64(packageID),
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
		"Message": "Packages fetched successfully",
		"Data":    response,
	})
}

// ViewPackages fetches packages for a coordinator.
func ViewPackages(ctx *gin.Context, client cpb.CoordinatorClient) {
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	_, id, err := middleware.ValidateToken(ctx, "coordinator")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
	}
	ID, _ := strconv.Atoi(id)
	ctxt := context.Background()
	response, err := client.CoordinatorViewPackages(ctxt, &cpb.View{
		ID:   int64(ID),
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
		"Message": "Packages fetched successfully",
		"Data":    response,
	})
}

// CoordinatorViewCategory fetches categories.
func CoordinatorViewCategory(ctx *gin.Context, client cpb.CoordinatorClient) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	ctxt := context.Background()
	response, err := client.ViewCategories(ctxt, &cpb.View{
		Page: int64(page),
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
		"Message": "Categories fetched successfully",
		"Data":    response,
	})
}
