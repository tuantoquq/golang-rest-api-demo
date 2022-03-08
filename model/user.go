package model

//define User for REST
type User struct {
	ID        int    `json:"id" binding:"required"`
	Email     string `json:"user_email" binding:"required,email" gorm:"unique"`
	Password  string `json:"user_password" binding:"required"`
	FirstName string `json:"user_firstname"`
	LastName  string `json:"user_lastname"`
	Phone     string `json:"user_phone"`
	Address   string `json:"user_address"`
}
