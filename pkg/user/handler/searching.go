package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

// validateQueryParam validates the query parameter and returns its value.
func validateQueryParam(ctx *gin.Context, paramName string) (string, bool) {
	value := ctx.GetHeader(paramName)
	if value == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  fmt.Sprintf("%s missing", paramName),
		})
		return "", false
	}
	return value, true
}

// SearchPackage searches for packages based on given parameters.
func SearchPackage(ctx *gin.Context, client pb.UserClient) {
	catagoryIDStr, valid1 := validateQueryParam(ctx, "catagoryid")
	pickupPlace, valid2 := validateQueryParam(ctx, "pickup_place")
	finalDestination, valid3 := validateQueryParam(ctx, "finaldestination")
	date, valid4 := validateQueryParam(ctx, "fromdate")
	if !valid1 || !valid2 || !valid3 || !valid4 {
		return
	}
	travelerCountStr := ctx.DefaultQuery("traveler_count", "1")
	pageStr := ctx.DefaultQuery("page", "1")
	endDate := ctx.DefaultQuery("enddate", "")
	maxDestination := ctx.DefaultQuery("maxdestination", "3")
	maxStops, _ := strconv.Atoi(maxDestination)
	destinations := ctx.QueryArray("destination")

	catagoryID, _ := strconv.Atoi(catagoryIDStr)
	travelerCount, _ := strconv.Atoi(travelerCountStr)
	page, _ := strconv.Atoi(pageStr)

	ctxt := context.Background()
	response, err := client.UserSearchPackage(ctxt, &pb.UserSearch{
		Category_ID:       int64(catagoryID),
		Traveler_Count:    int64(travelerCount),
		Pickup_Place:      pickupPlace,
		Final_Destination: finalDestination,
		Date:              date,
		Page:              int64(page),
		End_Date:          endDate,
		Max_Destination:   int64(maxStops),
		Destination:       destinations,
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
