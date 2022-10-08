package user

type User struct {
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
}
