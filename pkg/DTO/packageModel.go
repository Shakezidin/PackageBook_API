package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONB type for handling JSON data in the database
type JSONB []interface{}

// Value used to retrive value
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan helps to scan values
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

type Addpackage struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	StartLocation    string `json:"startlocation" validate:"required"`
	StartDate        string `json:"startdate" validate:"required"`
	StartTime        string `json:"starttime"`
	EndDate          string `json:"enddate" validate:"required"`
	Price            int    `json:"price" validate:"required"`
	Image            string `json:"image" validate:"required"`
	DestinationCount int    `json:"destinationcount" validate:"required"`
	Destination      string `json:"destination" validate:"required"`
	MaxCapacity      int64  `json:"maxcapacity" validate:"required"`
}

type AddDestination struct {
	DestinationName    string `json:"destinationname" validate:"required"`
	Description        string `json:"description" validate:"required"`
	Image              string `json:"image" validate:"required"`
	TransportationMode string `json:"transportationmode"`
	ArrivalLocation    string `json:"arrivallocation"`
}

type AddActivities struct {
	ActivityName string `json:"activityname" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Price        int    `json:"price" validate:"required"`
	Duration     string `json:"duration" validate:"required"`
	Location     string `json:"location" validate:"required"`
	ActivityType string `json:"activitytype" validate:"required"`
	Image        string `json:"image" validate:"required"`
	Date         string `json:"date" validate:"required"`
	Time         string `json:"time" validate:"required"`
}

type AddPromotion struct {
	Description    string `json:"description" validate:"required"`
	Discount       string `json:"discount" validate:"required"`
	StartDate      string `json:"startdate" validate:"required"`
	EndDate        string `json:"enddate" validate:"required"`
	Image          string `json:"image" validate:"required"`
	PromotionLevel string `josn:"promotionlevel" validate:"required"`
}

type AddCategory struct {
	Category string `json:"catagory" validate:"required"`
}

type FoodMenu struct {
	ID        int    `json:"id"`
	PackageID int    `json:"package_id" `
	Breakfast string `json:"breakfast" validate:"required"`
	Lunch     string `json:"lunch" validate:"required"`
	Dinner    string `json:"dinner" validate:"required"`
	Date      string `json:"date"   validate:"required"`
}
