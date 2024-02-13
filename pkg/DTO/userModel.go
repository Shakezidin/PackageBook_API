package dto

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Booking struct {
	Name          string `json:"name"`
	Age           int    `json:"age"`
	Address       string `json:"address"`
	Phone         int32  `json:"phone"`
	Adhaar        int    `json:"adhaar"`
	PackageId     int
	DestinationId int
	ActivityId    int
}

type Bookings struct {
	Users []Booking
}
