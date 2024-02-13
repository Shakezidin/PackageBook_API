package dto

// User represents the structure for user data.
type User struct {
	Name     string `json:"name" validate:"required"`     // Name of the user.
	Email    string `json:"email" validate:"required"`    // Email of the user.
	Phone    string `json:"phone" validate:"required"`    // Phone number of the user.
	Password string `json:"password" validate:"required"` // Password of the user.
}

// Booking represents the structure for a booking.
type Booking struct {
	Name          string `json:"name"`    // Name of the person booking.
	Age           int    `json:"age"`     // Age of the person booking.
	Address       string `json:"address"` // Address of the person booking.
	Phone         int32  `json:"phone"`   // Phone number of the person booking.
	Adhaar        int    `json:"adhaar"`  // Aadhar card number of the person booking.
	PackageId     int    // ID of the booked package.
	DestinationId int    // ID of the booked destination.
	ActivityId    int    // ID of the booked activity.
}

// Bookings represents a collection of bookings.
type Bookings struct {
	Users []Booking // List of bookings.
}
