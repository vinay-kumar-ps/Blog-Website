package util

import (
	"time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v4"
	// "golang.org/x/oauth2/jwt"
)

const secretkey = "secret"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	return claims.SignedString([]byte(secretkey))
}
func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, jwt.StandardClaims{}, func(t *jwt.Token)(interface{},error){
		return []byte(secretkey),nil
	})
	if err != nil||! token.Valid {
		return "",err
	}
	 claims :=token.Claims.(*jwt.StandardClaims)
	 return claims.Issuer,nil
	
}
