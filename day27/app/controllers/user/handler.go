package user

import (
	"crypto/sha1"
	"fmt"
	User "it/day27/app/models/user"
)

// 註冊使用者
// 這邊可以重複註冊，如果不希望有重複的 email 出現，可在資料庫中設定 email 屬性為 unique
func RegisterUser(email, password string) bool {
	passwordSha1 := sha1It(password)
	err := User.CreateUser(email, passwordSha1)
	return err == nil
}

// 使用者登入
// 如果資料庫中沒有找到對應的使用者帳密，回傳 err = record not found，有找到則 err = nil
func LoginUser(email, password string) bool {
	passwordSha1 := sha1It(password)
	err := User.LoginUser(email, passwordSha1)
	return err == nil
}

// 加密字串
const EncryptCode = "it"

// sha加密
func sha1It(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	bs := h.Sum([]byte(EncryptCode))
	encryptCode := fmt.Sprintf("%x", bs)
	return encryptCode
}
