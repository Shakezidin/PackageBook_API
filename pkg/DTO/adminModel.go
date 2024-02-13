package dto

// Admin represents the structure for administrative user credentials.
type Admin struct {
	Username string `json:"username" validate:"required"` // Username represents the administrator's username.
	Password string `json:"password" validate:"required"` // Password represents the administrator's password.
}
