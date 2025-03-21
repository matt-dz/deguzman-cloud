package handlers

type HandleLoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type HandleSignupBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type HandleEmailVerificationBody struct {
	Code string `json:"code"`
}

type HandleLoginResponse struct {
	Redirect string `json:"redirect"`
}
