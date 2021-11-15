package domain

type Music struct {
	Id 	 int  `json:"id"`
	Title string `json:"title"`
	Duration int `json:"duration"`
	Singer string `json:"singer"`
	CreatedOn string `json:"created_on"`
}