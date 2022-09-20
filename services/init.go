package services

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var JwtService *jwtService

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	JwtService = &jwtService{}
	JwtService.JwtSecret = os.Getenv("JWT_SECRET")
	t := os.Getenv("JWT_EXPIRES")
	rt := os.Getenv("REFRESH_TOKEN_EXPIRES")
	p, err := strconv.Atoi(t)
	if err != nil {
		panic(err)
	}
	rp, err := strconv.Atoi(rt)
	if err != nil {
		panic(err)
	}
	JwtService.JwtExpireTime = p
	JwtService.RefreshTokenExpireTime = rp

}
