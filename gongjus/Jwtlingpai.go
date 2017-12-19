package gongjus

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const (
	Jwt_mima = "clz"
	Jwt_yan  = "abc"
)

type Yanzhengclaims struct {
	Zhanghao string
	jwt.StandardClaims
}

func Shengchenglingpai(user string) (string, error) {
	key := []byte(Jwt_mima + Jwt_yan)
	claims := Yanzhengclaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 11).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
func Yanzhenglingpai(lingpai string) *Yanzhengclaims {
	token, err := jwt.ParseWithClaims(
		lingpai,
		&Yanzhengclaims{},
		func(token *jwt.Token) (interface{}, error) {
			key := []byte(Jwt_mima + Jwt_yan)
			return key, nil
		},
	)
	if err != nil {
		log.Println("解析token失败", err)
		return nil
	}
	if claims, ok := token.Claims.(*Yanzhengclaims); ok && token.Valid {
		return claims
	}
	return nil
}
