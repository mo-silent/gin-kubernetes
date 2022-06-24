package service

import (
	"errors"
	"time"

	"gitee.com/MoGD/gin-kubernetes/global"
	"github.com/golang-jwt/jwt/v4"
)

//自定义payload结构体,不建议直接使用 dgrijalva/jwt-go `jwt.StandardClaims`结构体.因为他的payload包含的用户信息太少.
type UserClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (j *UserClaims) CreateToken(username string) (string, error) {

	// Create claims while leaving out some of the optional fields
	claims := UserClaims{
		username,
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			NotBefore: jwt.NewNumericDate(time.Unix(time.Now().Unix()-1000, 0)),                                     // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.CONFIG.JWT.ExpiresTime) * time.Hour)), // 过期时间 1小时  配置文件
			Issuer:    global.CONFIG.JWT.Issuer,                                                                     // 签名的发行者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.CONFIG.JWT.SigningKey))
}

func (j *UserClaims) ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.CONFIG.JWT.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("couldn't handle this token")
	}
}

func NewClaims() *UserClaims {
	return &UserClaims{}
}
