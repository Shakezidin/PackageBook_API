package dto

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}

type Booking struct {
	Name          string `json:"name"`
	Age           int    `json:"age"`
	Address       string `json:"address"`
	Phone         int    `json:"phone"`
	Adhaar        int    `json:"adhaar"`
	PackageId     int
	DestinationId int
	ActivityId    int
}

type Bookings struct {
	Users []Booking
}
