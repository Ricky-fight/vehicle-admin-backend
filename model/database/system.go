package database

type User struct {
	Account  string `gorm:"unique"`
	Password string
	Model
}
