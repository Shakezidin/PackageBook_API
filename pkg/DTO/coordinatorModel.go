package dto

type Addpackage struct {
	Name string `json:"name"`
	StartLocation    string `json:"startlocation"`
	EndLocation      string `json:"endlocation"`
	StartDateTime    string `json:"startdatetime"`
	EndDateTime      string `json:"enddatetiime"`
	Price            int    `json:"price"`
	Image            string `json:"image"`
	DestinationCount int    `json:"destinationcount"`
	Destinations     string `json:"destinations"`
}

type AddDestination struct {
	DestinationName string `json:"destinationname"`
	Description     string `json:"description"`
	MinPrice        int    `json:"minPrice"`
	Image           string `json:"image"`
	MaxCapacity     int    `json:"maxcapacity"`
}

type AddActivities struct {
	ActivityName   string `json:"activityname"`
	Description    string `json:"description"`
	Price          int    `json:"price"`
	Duration       string `json:"duration"`
	Location       string `json:"location"`
	DefficultLevel string `json:"difficultylevel"`
	Image          string `json:"image"`
}

type AddPromotion struct {
	Description    string `json:"description"`
	Discount       string `json:"discount"`
	StartDate      string `json:"startdate"`
	EndDate        string `json:"enddate"`
	Image          string `json:"image"`
	PromotionLevel string `josn:"promotionlevel"`
}

