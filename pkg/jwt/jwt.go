package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	//"github.com/golang-jwt/jwt"
	"time"
)

//var jwtSecret = []byte(settings.Conf.JwtSecret)
var MySecret = []byte("夏天夏天悄悄过去11")

const TokenExpireDuration = time.Hour * 2

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(userID int64, username string) (atoken, rtoken string, err error) {

	claims := Claims{
		userID,
		username,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(),
			Issuer:    "rolin",
		},
	}
	fmt.Println(claims)

	//tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	atoken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(MySecret)
	rtoken, err = jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		Issuer:    "rolin",
	}).SignedString(MySecret)
	return
}
func ParseToken(token string) (*Claims, error) {
	//解析token
	var mc = new(Claims)
	tokenClaims, err := jwt.ParseWithClaims(token, mc, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")

}

//刷新token
//func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
//	//refresh token 无效直接返回
//	if err, _ := jwt.Parse(rToken, KeyFunc); err != nil {
//		return
//	}
//	//从旧的token解析出claims
//	var claim Claims
//	_, err = jwt.ParseWithClaims(aToken, &claim, KeyFunc)
//	v, _ := err.(jwt.ValidationError)
//	//如果是过期错误，且rtoken没有过期，就更新token
//	if v.Errors == jwt.ValidationErrorExpired {
//		return GenerateToken(claim.UserID, claim.Username)
//	}
//	return
//
//}

func KeyFunc(token *jwt.Token) (interface{}, error) {
	return nil, nil
}
