package domain

type User struct {
	Id 	 int  `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Gender string `json:"gender"`
	CreatedOn string `json:"created_on"`
}