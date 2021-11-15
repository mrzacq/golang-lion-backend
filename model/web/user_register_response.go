package web

type UserRegisterResponse struct {
	Id        int    `validate:"required,min=1" json:"id"`
	FullName string `validate:"required,min=1,max=255" json:"full_name"`
	Username string `validate:"required,min=1,max=255" json:"username"`
	Gender string `validate:"required,min=1,max=255" json:"gender"`
	CreatedOn string `validate:"required" json:"created_on"`
}