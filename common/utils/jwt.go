package utils

import (
	"errors"
	"github.com/astaxie/beego"

	"log"
	"time"
	"youblog/modules"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	key []byte = []byte("-jwt-a1624000875@163.com")
)

//生成json web token
func GenerateToken(user *modules.SysUser) string {
	//有效载荷
	/*	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Add(time.Hour*24*7).Unix()),	//设置过期时间30天
		Issuer:    "daimao",	//签发人
	}*/
	claims := jwt.MapClaims{
		"Id":        user.UserId,
		"username":  user.Username,
		"NotBefore": int64(time.Now().Unix()),
		"ExpiresAt": int64(time.Now().Add(time.Hour * 24 * 7).Unix()), //设置过期时间30天
		"Issuer":    "daimao",                                         //签发人
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //加密
	signingString, err := token.SignedString(key)              //签名
	if err != nil {
		beego.Error("token签名失败:", err)
	}
	return signingString
}

// 校验token是否有效
func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		log.Println("警告有人携带恶意cookie.", err)
		return false
	}

	return true
}

//解析token,从token中获取userId
func GetUserIDFromToken(userToken string) (userId int64, err error) {
	token, err := jwt.Parse(userToken, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("无法将Claims转化为MapClaims")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token令牌无效")
		return
	}
	id := claims["Id"].(float64)
	return int64(id), nil
}
