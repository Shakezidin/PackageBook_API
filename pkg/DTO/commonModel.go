package dto

// Login represents the structure for user login.
type Login struct {
	Email    string `json:"email" validate:"required"` // Email represents the user's email address.
	Password string `json:"password" validate:"required"` // Password represents the user's password.
}

// Otp represents the structure for OTP verification.
type Otp struct {
	OTP   int    `json:"otp" validate:"required"` // OTP represents the one-time password for verification.
	Email string `json:"email" validate:"required"` // Email represents the user's email address.
}

// PasswordChange represents the structure for changing passwords.
type PasswordChange struct {
	OldPassword string `json:"oldpassword" validate:"required"` // OldPassword represents the user's old password.
	NewPassword string `json:"newpassword" validate:"required"` // NewPassword represents the user's new password.
}

// ForgetPassword represents the structure for initiating the forget password process.
type ForgetPassword struct {
	Phone string `json:"phone" validate:"required"` // Phone represents the user's phone number for password recovery.
}

// NewChange represents the structure for confirming a new password.
type NewChange struct {
	NewPassword     string `json:"newpassword" validate:"required"` // NewPassword represents the user's new password.
	ConfirmPassword string `json:"confirmpassword" validate:"required"` // ConfirmPassword represents the confirmation of the new password.
}

// PswrdOtp represents the structure for OTP verification during password reset.
type PswrdOtp struct {
	OTP   string `json:"otp" validate:"required"` // OTP represents the one-time password for verification.
	Phone string `json:"phone" validate:"required"` // Phone represents the user's phone number.
}
