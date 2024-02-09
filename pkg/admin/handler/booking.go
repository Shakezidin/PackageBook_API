package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

func ViewBookings(ctx *gin.Context, client pb.AdminClient) {
	pkgId := ctx.GetHeader("id")
	payment := ctx.DefaultQuery("payment", "full")
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	id, _ := strconv.Atoi(pkgId)

	ctxt := context.Background()
	response, err := client.AdminViewBookings(ctxt, &pb.AdminView{
		Id:     int64(id),
		Status: payment,
		Page:   int64(page),
	})

	if err != nil {
		errMsg := "error while fetching bookings"
		log.Println(errMsg, err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errMsg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "bookings fetched successfully",
		"data":    response,
	})
}

func ViewBooking(ctx *gin.Context, client pb.AdminClient) {
	id := ctx.GetHeader("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		errMsg := "booking ID missing or invalid"
		log.Println(errMsg, err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errMsg,
		})
		return
	}

	ctxt := context.Background()
	response, err := client.AdminViewBooking(ctxt, &pb.AdminView{
		Id: int64(ID),
	})

	if err != nil {
		errMsg := "error while fetching booking"
		log.Println(errMsg, err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errMsg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "booking fetched successfully",
		"data":    response,
	})
}

// BookingSearchCriteria represents the search criteria for bookings.
type BookingSearchCriteria struct {
	PaymentMode     string `json:"paymentMode,omitempty"`
	BookingStatus   string `json:"bookingStatus,omitempty"`
	CancelledStatus bool   `json:"cancelledStatus,omitempty"`
	UserEmail       string `json:"userEmail,omitempty"`
	BookingID       string `json:"bookingID,omitempty"`
	BookDate        string `json:"bookDateFrom,omitempty"`
	StartDate       string `json:"startDateFrom,omitempty"`
	CoordinatorID   uint   `json:"coordinatorID,omitempty"`
	CatageryId      uint   `json:"categoryid,omitempty"`
}

func isValidDateFormat(dateStr string) bool {
	_, err := time.Parse("02-01-2006", dateStr)
	return err == nil
}

func FilterBookings(ctx *gin.Context, client pb.AdminClient) {
	var Searching BookingSearchCriteria

	if err := ctx.BindJSON(&Searching); err != nil {
		log.Printf("error binding JSON")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
			"msg":    "error",
		})
		return
	}

	if Searching.BookDate != "" && !isValidDateFormat(Searching.BookDate) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid bookDate format. Please use the format DD-MM-YYYY.",
		})
		return
	}

	if Searching.StartDate != "" && !isValidDateFormat(Searching.StartDate) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "Invalid startDate format. Please use the format DD-MM-YYYY.",
		})
		return
	}
	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	var ctxt = context.Background()
	response, err := client.AdminSearchBooking(ctxt, &pb.AdminBookingSearchCriteria{
		PaymentMode:     Searching.PaymentMode,
		BookingStatus:   Searching.BookingStatus,
		CancelledStatus: Searching.CancelledStatus,
		UserEmail:       Searching.UserEmail,
		BookingId:       Searching.BookingID,
		BookDate:        Searching.BookDate,
		StartDate:       Searching.StartDate,
		CoordinatorId:   uint32(Searching.CoordinatorID),
		Page:            int64(pageInt),
		CatageryId:      int64(Searching.CatageryId),
	})

	if err != nil {
		log.Printf("bookings fetching  error", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("bookings fetched succesfully"),
		"data":    response,
	})
}
