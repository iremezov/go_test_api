package models

type User struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Login string `json:"login"`
	Email string `json:"email"`
	State int    `json:"state"`
}
