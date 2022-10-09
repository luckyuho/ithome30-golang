package user

type User struct {
	Id       int    `gorm:"id"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
}
