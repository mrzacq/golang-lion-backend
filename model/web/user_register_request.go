package web

type UserRegisterRequest struct {
	FullName string `validate:"required,min=2,max=255" json:"full_name"`
	Username string `validate:"required,min=2,max=255" json:"username"`
	Password string `validate:"required,min=8,max=255" json:"password"`
	Gender string `validate:"required,min=2,max=255" json:"gender"`
}