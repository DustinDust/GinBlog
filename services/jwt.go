package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

type jwtService struct {
	JwtSecret              string
	JwtExpireTime          int
	RefreshTokenExpireTime int
}

func (j *jwtService) ParseJwtFromBearerToken(bearerTokenString string) (*JwtClaims, error) {
	if !(strings.HasPrefix(bearerTokenString, "Bearer ")) {
		return nil, errors.New("Not a valid bearer token")
	}
	tokenString := strings.Fields(bearerTokenString)[1]
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.JwtSecret), nil
	})
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (j *jwtService) GenerateJwt(claims JwtClaims) (string, error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(j.JwtExpireTime) * time.Second))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.JwtSecret))
	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func (j *jwtService) GenerateRefreshJwt(claims JwtClaims) (string, error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(j.RefreshTokenExpireTime) * time.Second))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.JwtSecret))
	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}
