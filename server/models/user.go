package models

type User struct {
	ID 				int 		`json:"id"`
	Name 			string	`json:"name"`
	Email 		string 	`json:"email"`//in the tutorial, set this as `gorm:"unique"`
	Password 	[]byte	`json:"password"`
}