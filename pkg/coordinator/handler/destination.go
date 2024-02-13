package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	dto "github.com/Shakezidin/pkg/DTO"
	cpb "github.com/Shakezidin/pkg/coordinator/pb"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// AddDestination handles the addition of a new destination.
func AddDestination(ctx *gin.Context, client cpb.CoordinatorClient) {
	var destination dto.AddDestination

	packageIDStr := ctx.GetHeader("id")
	packageID, err := strconv.Atoi(packageIDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Package ID missing",
		})
		return
	}

	if err := ctx.BindJSON(&destination); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Error binding JSON",
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(destination)
	if err != nil {
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

	var ctxt = context.Background()
	response, err := client.CoordinatorAddDestination(ctxt, &cpb.Destination{
		DestinationName:    destination.DestinationName,
		Description:        destination.Description,
		PackageID:          int64(packageID),
		Image:              destination.Image,
		TransportationMode: destination.TransportationMode,
		ArrivalLocation:    destination.ArrivalLocation,
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
		"Message": fmt.Sprintf("%v destination created successfully", destination.DestinationName),
		"Data":    response,
	})
}

// ViewDestination fetches destination information based on package ID.
func ViewDestination(ctx *gin.Context, client cpb.CoordinatorClient) {
	packageIDStr := ctx.GetHeader("id")
	packageID, err := strconv.Atoi(packageIDStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Package ID missing",
		})
		return
	}

	var ctxt = context.Background()
	response, err := client.CoordinatorViewDestination(ctxt, &cpb.View{
		Id: int64(packageID),
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
		"Message": "Destination fetched successfully",
		"Data":    response,
	})
}

const mapboxAccessToken = "pk.eyJ1Ijoic2ludXppZGluIiwiYSI6ImNscmdpdXh0bjBod2wyam81dGt1dHppN28ifQ.rKMd949jNDCKZr1jC2qfeA"

// SuggestLocation provides location suggestions using Mapbox Geocoding API.
func SuggestLocation(c *gin.Context) {
	// Get location from header
	location := c.GetHeader("location")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Location header is required",
		})
		return
	}

	// Construct Mapbox Geocoding API URL
	apiURL := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?access_token=%s", location, mapboxAccessToken)

	// Make HTTP request to Mapbox Geocoding API
	response, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}
	defer response.Body.Close()

	// Read and parse the API response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Parse JSON response
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	// Extract main data from the response
	var suggestions []string
	features, ok := result["features"].([]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  "Invalid response format",
		})
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
	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Location suggestions fetched successfully",
		"Data":    suggestions,
	})
}
