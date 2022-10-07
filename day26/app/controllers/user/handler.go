package user

import User "it/day26/app/models/user"

// 註冊使用者
// 這邊可以重複註冊，如果不希望有重複的 email 出現，可在資料庫中設定 email 屬性為 unique
func RegisterUser(email, password string) bool {
	err := User.CreateUser(email, password)
	return err == nil
}

// 使用者登入
// 如果資料庫中沒有找到對應的使用者帳密，回傳 err = record not found，有找到則 err = nil
func LoginUser(email, password string) bool {
	err := User.LoginUser(email, password)
	return err == nil
}
