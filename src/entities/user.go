package entities

type User struct {
	ID uint	`json: id gorm:"primary_key"`
    FirstName string `json:"first_name"`
	Email  string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
