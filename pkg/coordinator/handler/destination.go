package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AddDestination(ctx *gin.Context, client cpb.CoordinatorClient) {
	var destination dto.AddDestination

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

	if err := ctx.BindJSON(&destination); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error binding JSON",
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(destination)
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

	var ctxt = context.Background()
	response, err := client.CoordinatorAddDestination(ctxt, &cpb.Destination{
		DestinationName: destination.DestinationName,
		Description:     destination.Description,
		PackageID:       int64(packageId),
		Image:           destination.Image,
	})

	if err != nil {
		log.Printf("destination %s creattion error", destination.DestinationName, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v destination created succesfully", destination.DestinationName),
		"data":    response,
	})
}

func ViewDestination(ctx *gin.Context, client cpb.CoordinatorClient) {
	packageIdStr := ctx.GetHeader("id")
	packageId, err := strconv.Atoi(packageIdStr)
	if err != nil {
		fmt.Println("destination missing")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.CoordinatorViewDestination(ctxt, &cpb.View{
		Id: int64(packageId),
	})

	if err != nil {
		log.Printf("destination fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("destination fetched succesfully"),
		"data":    response,
	})
}

const mapboxAccessToken = "pk.eyJ1Ijoic2ludXppZGluIiwiYSI6ImNscmdpdXh0bjBod2wyam81dGt1dHppN28ifQ.rKMd949jNDCKZr1jC2qfeA"

func SuggestLocation(c *gin.Context) {
	// Get location from header
	location := c.GetHeader("location")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location header is required"})
		return
	}

	// Construct Mapbox Geocoding API URL
	apiURL := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?access_token=%s", location, mapboxAccessToken)

	// Make HTTP request to Mapbox Geocoding API
	response, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error: %s", err)})
		return
	}
	defer response.Body.Close()

	// Read and parse the API response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error: %s", err)})
		return
	}

	// Parse JSON response
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error: %s", err)})
		return
	}

	// Extract main data from the response
	var suggestions []string
	features, ok := result["features"].([]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response format"})
		return
	}

	for _, feature := range features {
		featureData, ok := feature.(map[string]interface{})
		if !ok {
			continue
		}

		name, ok := featureData["place_name"].(string)
		if !ok {
			continue
		}

		suggestions = append(suggestions, name)
	}

	// Display suggestions to the user
	c.JSON(http.StatusOK, suggestions)
}
