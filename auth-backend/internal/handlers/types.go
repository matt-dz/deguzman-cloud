package handlers

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type EmailVerificationBody struct {
	Code string `json:"code"`
}

type AuthBody struct {
	Role string `json:"role"`
}

type LoginResponse struct {
	Redirect string `json:"redirect"`
}
