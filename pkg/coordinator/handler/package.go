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

func AddPackage(ctx *gin.Context, client cpb.CoordinatorClient) {
	var pkg dto.Addpackage

	destination := ctx.GetHeader("destination")
	startlocation := ctx.GetHeader("startlocation")
	categoryIdStr := ctx.GetHeader("id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		fmt.Println("categoryID missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	_, Id, err := middleware.ValidateToken(ctx, "coordinator")
	if err != nil {
		log.Printf("token validation error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	if err := ctx.BindJSON(&pkg); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}
	pkg.Destination = destination
	pkg.StartLocation = startlocation

	validate := validator.New()
	err = validate.Struct(pkg)
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

	id, _ := strconv.Atoi(Id)

	startdate, err := time.Parse("02-01-2006", pkg.StartDate)
	_, err = time.Parse("03:04 PM", pkg.StartTime)
	enddate, err := time.Parse("02-01-2006", pkg.EndDate)
	if err != nil {
		log.Printf("date fromat error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	if !enddate.Add(24 * time.Hour).After(startdate) {
		log.Printf("error in start date and enddate")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}
	var ctxt = context.Background()
	response, err := client.CoordinatorAddPackage(ctxt, &cpb.Package{
		CoorinatorId:     int64(id),
		Packagename:      pkg.Name,
		Startdate:        pkg.StartDate,
		Starttime:        pkg.StartTime,
		Startlocation:    pkg.StartLocation,
		Enddate:          pkg.EndDate,
		Price:            int64(pkg.Price),
		Image:            pkg.Image,
		DestinationCount: int64(pkg.DestinationCount),
		Destination:      pkg.Destination,
		CategoryId:       int64(categoryId),
		Description:      pkg.Description,
		MaxCapacity:      pkg.MaxCapacity,
	})

	if err != nil {
		log.Printf("pakcage %s creattion error", pkg.Name, err.Error())
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

func ViewPackage(ctx *gin.Context, client cpb.CoordinatorClient) {
	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		fmt.Println("packageId missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.CoordinatorViewPackage(ctxt, &cpb.View{
		Id: int64(packageId),
	})

	if err != nil {
		log.Printf("package fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("packages fetched succesfully"),
		"data":    response,
	})

}

func ViewPackages(ctx *gin.Context, client cpb.CoordinatorClient) {
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	_, id, err := middleware.ValidateToken(ctx, "coordinator")
	if err != nil {
		log.Printf("token validation error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	}
	Id, _ := strconv.Atoi(id)
	var ctxt = context.Background()
	response, err := client.CoordinatorViewPackages(ctxt, &cpb.View{
		Id:   int64(Id),
		Page: int64(pageInt),
	})

	if err != nil {
		log.Printf("packages fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("packages fetched succesfully"),
		"data":    response,
	})
}

func CoordinatorViewCatagory(ctx *gin.Context, client cpb.CoordinatorClient) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	var ctxt = context.Background()
	response, err := client.ViewCatagories(ctxt, &cpb.View{
		Page: int64(page),
	})

	if err != nil {
		log.Printf("catagories fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("catagories fetched succesfully"),
		"data":    response,
	})
}
