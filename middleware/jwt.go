package middleware

/**
* @Author: Xenolies
* @Date: 2023/4/16 19:51
* @Version: 1.0
 */

import (
	"bolg_go/config"
	"bolg_go/model"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type Claims struct {
	ID       string
	UserName string
	Password string

	jwt.StandardClaims
}

var Secret = []byte(config.GlobalConfig.Admin.JwtSecret) // 定义secret

var ExpirationTime, _ = strconv.Atoi(config.GlobalConfig.Admin.ExpirationTime)

// MakeToken 生成 Token
// userName, userPassword, strconv.FormatInt(sfID, 10)
// userName string, userPwd string, ID string
func MakeToken(users model.Users) (tokenString string, err error) {
	timeNow := time.Now()
	expireTime := timeNow.Add(time.Duration(ExpirationTime) * time.Minute) // 使用用户设置的token过期时间
	claim := Claims{
		ID:       strconv.FormatInt(users.ID, 10),
		UserName: users.UserName,
		Password: users.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			IssuedAt:  timeNow.Unix(),    // 签发时间
			NotBefore: timeNow.Unix(),    // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err = token.SignedString(Secret)
	return tokenString, err
}

// ParseToken Token解码
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
