package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	//"github.com/golang-jwt/jwt"
	"time"
)

//var jwtSecret = []byte(settings.Conf.JwtSecret)
var MySecret = []byte("夏天夏天悄悄过去")

const TokenExpireDuration = time.Hour * 2

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(userID int64, username string) (string, error) {

	claims := Claims{
		userID,
		username,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "rolin",
		},
	}
	fmt.Println(claims)

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(MySecret)
	return token, err
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
