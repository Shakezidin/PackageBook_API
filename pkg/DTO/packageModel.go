package dto

// Addpackage represents the structure for adding a tour package.
type Addpackage struct {
	Name             string `json:"name"`                // Name of the tour package.
	Description      string `json:"description"`         // Description of the tour package.
	StartLocation    string `json:"startlocation"`       // Starting location of the tour package.
	StartDate        string `json:"startdate"`           // Start date of the tour package.
	StartTime        string `json:"starttime,omitempty"` // Start time of the tour package.
	EndDate          string `json:"enddate"`             // End date of the tour package.
	Price            int    `json:"price"`               // Price of the tour package.
	Image            string `json:"image"`               // Image representing the tour package.
	DestinationCount int    `json:"destinationcount"`    // Count of destinations included in the tour package.
	Destination      string `json:"destination"`         // Main destination of the tour package.
	MaxCapacity      int64  `json:"maxcapacity"`         // Maximum capacity of the tour package.
}

// AddDestination represents the structure for adding a destination to a tour package.
type AddDestination struct {
	DestinationName    string `json:"destinationname"`    // Name of the destination.
	Description        string `json:"description"`        // Description of the destination.
	Image              string `json:"image"`              // Image representing the destination.
	TransportationMode string `json:"transportationmode"` // Transportation mode to reach the destination.
	ArrivalLocation    string `json:"arrivallocation"`    // Location where the destination is reached.
}

// AddActivities represents the structure for adding activities to a tour package.
type AddActivities struct {
	ActivityName string `json:"activityname"` // Name of the activity.
	Description  string `json:"description"`  // Description of the activity.
	Price        int    `json:"price"`        // Price of the activity.
	Duration     string `json:"duration"`     // Duration of the activity.
	Location     string `json:"location"`     // Location where the activity takes place.
	ActivityType string `json:"activitytype"` // Type of the activity.
	Image        string `json:"image"`        // Image representing the activity.
	Date         string `json:"date"`         // Date when the activity is scheduled.
	Time         string `json:"time"`         // Time when the activity starts.
}

// AddPromotion represents the structure for adding a promotion to a tour package.
type AddPromotion struct {
	Description    string `json:"description"`    // Description of the promotion.
	Discount       string `json:"discount"`       // Discount offered by the promotion.
	StartDate      string `json:"startdate"`      // Start date of the promotion.
	EndDate        string `json:"enddate"`        // End date of the promotion.
	Image          string `json:"image"`          // Image representing the promotion.
	PromotionLevel string `json:"promotionlevel"` // Level or type of the promotion.
}

// AddCategory represents the structure for adding a category to the system.
type AddCategory struct {
	Category string `json:"catagory"` // Category name.
}

// FoodMenu represents the structure for defining food options for a tour package.
type FoodMenu struct {
	ID        int    `json:"id"`         // Unique identifier for the food menu item.
	PackageID int    `json:"package_id"` // ID of the tour package associated with the food menu.
	Breakfast string `json:"breakfast"`  // Breakfast option for the package.
	Lunch     string `json:"lunch"`      // Lunch option for the package.
	Dinner    string `json:"dinner"`     // Dinner option for the package.
	Date      string `json:"date"`       // Date for which the food options are available.
}
