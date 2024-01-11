package dto

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Otp struct {
	OTP   int    `json:"otp" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type PasswordChange struct {
	OldPassword string `json:"oldpassword" validate:"required"`
	NewPassword string `json:"newpassword" validate:"required"`
}

type ForgetPassword struct {
	Email string `json:"email" validate:"required"`
}

type NewChange struct {
	NewPassword     string `json:"newpassword" validate:"required"`
	ConfirmPassword string `json:"confirmpassword" validate:"required"`
}
