package user

import (
	"fmt"
	"it/day27/database"
	"it/day27/global"
)

// 新增資料
func CreateUser(email, password string) error {
	user := User{Email: email, Password: password}
	err := database.Get().Table(global.UserTable).Create(&user).Error
	return err
}

// 從資料庫找有沒有相對應的 email 與 password
func LoginUser(email, password string) error {
	user := User{Email: email, Password: password}
	err := database.Get().Table(global.UserTable).Where("email = ? and password = ?", email, password).Scan(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}
