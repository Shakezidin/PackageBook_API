package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	pb "github.com/Shakezidin/pkg/user/userpb"
	"github.com/gin-gonic/gin"
)

func validateQueryParam(ctx *gin.Context, paramName string) (string, bool) {
	value := ctx.GetHeader(paramName)
	if value == "" {
		log.Printf("%s missing", paramName)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  fmt.Sprintf("%s missing", paramName),
		})
		return "", false
	}
	return value, true
}

func SearchPackage(ctx *gin.Context, client pb.UserClient) {
	catagoryIdStr, valid1 := validateQueryParam(ctx, "catagoryid")
	pickup_place, valid2 := validateQueryParam(ctx, "pickup_place")
	finaldestination, valid3 := validateQueryParam(ctx, "finaldestination")
	date, valid4 := validateQueryParam(ctx, "fromdate")
	if !valid1 || !valid2 || !valid3 || !valid4 {
		return
	}
	traveler_countStr := ctx.DefaultQuery("traveler_count", "1")
	pageStr := ctx.DefaultQuery("page", "1")
	endDate := ctx.DefaultQuery("enddate", "")
	maxDestination := ctx.DefaultQuery("maxdestination", "3")
	maxStops, _ := strconv.Atoi(maxDestination)
	var destinations = ctx.QueryArray("destination")

	catagoryId, _ := strconv.Atoi(catagoryIdStr)
	traveler_count, _ := strconv.Atoi(traveler_countStr)
	page, _ := strconv.Atoi(pageStr)

	var ctxt = context.Background()
	response, err := client.UserSearchPacakge(ctxt, &pb.UserSearch{
		CatagoryId:       int64(catagoryId),
		Travelercount:    int64(traveler_count),
		PickupPlace:      pickup_place,
		Finaldestination: finaldestination,
		Date:             date,
		Page:             int64(page),
		Enddate:          endDate,
		MaxDestination:   int64(maxStops),
		Destination:      destinations,
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
		"status":  http.StatusOK,
		"message": fmt.Sprintf("catagories fetched succesfully"),
		"data":    response,
	})

}
