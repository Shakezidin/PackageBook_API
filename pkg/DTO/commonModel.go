package dto

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Otp struct {
	OTP   int    `json:"otp"`
	Email string `json:"email"`
}

type PasswordChange struct {
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}

type ForgetPassword struct {
	Email string `json:"email"`
}

type NewChange struct {
	NewPassword     string `json:"newpassword"`
	ConfirmPassword string `json:"confirmpassword"`
}
