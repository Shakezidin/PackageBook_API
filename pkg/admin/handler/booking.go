package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	pb "github.com/Shakezidin/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

// ViewBookings handles the HTTP request to view bookings.
func ViewBookings(ctx *gin.Context, client pb.AdminClient) {
	pkgID := ctx.GetHeader("id")
	payment := ctx.DefaultQuery("payment", "full")
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	id, _ := strconv.Atoi(pkgID)

	ctxt := context.Background()
	response, err := client.AdminViewBookings(ctxt, &pb.AdminView{
		ID:     int64(id),
		Status: payment,
		Page:   int64(page),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Bookings fetched successfully",
		"Data":    response,
	})
}

// ViewBooking handles the HTTP request to view a booking.
func ViewBooking(ctx *gin.Context, client pb.AdminClient) {
	id := ctx.GetHeader("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		errMsg := "Booking ID missing or invalid"
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  errMsg,
		})
		return
	}

	ctxt := context.Background()
	response, err := client.AdminViewBooking(ctxt, &pb.AdminView{
		ID: int64(ID),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Status": http.StatusInternalServerError,
			"Error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Booking fetched successfully",
		"Data":    response,
	})
}

// BookingSearchCriteria represents the search criteria for bookings.
type BookingSearchCriteria struct {
	PaymentMode     string `json:"paymentmode,omitempty"`
	BookingStatus   string `json:"bookingstatus,omitempty"`
	CancelledStatus bool   `json:"cancelledstatus,omitempty"`
	UserEmail       string `json:"useremail,omitempty"`
	BookingID       string `json:"bookingid,omitempty"`
	BookDate        string `json:"bookdatefrom,omitempty"`
	StartDate       string `json:"startdatefrom,omitempty"`
	CoordinatorID   uint   `json:"coordinatorid,omitempty"`
	CategoryID      uint   `json:"categoryid,omitempty"`
}

func isValidDateFormat(dateStr string) bool {
	_, err := time.Parse("02-01-2006", dateStr)
	return err == nil
}

// FilterBookings handles the HTTP request to filter bookings based on criteria.
func FilterBookings(ctx *gin.Context, client pb.AdminClient) {
	var search BookingSearchCriteria

	if err := ctx.BindJSON(&search); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  err.Error(),
		})
		return
	}

	if search.BookDate != "" && !isValidDateFormat(search.BookDate) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Invalid bookdate format. Please use the format DD-MM-YYYY.",
		})
		return
	}

	if search.StartDate != "" && !isValidDateFormat(search.StartDate) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status": http.StatusBadRequest,
			"Error":  "Invalid startdate format. Please use the format DD-MM-YYYY.",
		})
		return
	}

	page := ctx.DefaultQuery("page", "1")
	pageInt, _ := strconv.Atoi(page)
	ctxt := context.Background()
	response, err := client.AdminSearchBooking(ctxt, &pb.AdminBookingSearchCriteria{
		Payment_Mode:     search.PaymentMode,
		Booking_Status:   search.BookingStatus,
		Cancelled_Status: search.CancelledStatus,
		User_Email:       search.UserEmail,
		Booking_ID:       search.BookingID,
		Book_Date:        search.BookDate,
		Start_Date:       search.StartDate,
		Coordinator_ID:   uint32(search.CoordinatorID),
		Page:            int64(pageInt),
		Category_ID:      int64(search.CategoryID),
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
		"Message": "Bookings fetched successfully",
		"Data":    response,
	})
}
