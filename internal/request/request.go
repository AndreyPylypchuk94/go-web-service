package request

type LoginRequest struct {
	Email    string `required:"true" json:"email"`
	Password string `required:"true" json:"password"`
}
