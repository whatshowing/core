package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MapClaims jwt.MapClaims

type JwtService interface {
	ExtractClaims(t string, secret string) (MapClaims, error)
	SignToken(mapClaims MapClaims, secret string, expSec time.Duration) (string, error)
	IsValid(t string, secret string) bool
	RefreshToken(
		refreshToken string,
		refreshSecret string,
		token string,
		tokenSecret string,
		expSec time.Duration,

	) (string, error)
}

type jwtService struct{}

func (j *jwtService) RefreshToken(
	refreshToken string,
	refreshSecret string,
	token string,
	tokenSecret string,
	expSec time.Duration,
) (string, error) {

	if ok := j.IsValid(refreshToken, refreshSecret); !ok {
		return "", errors.New("refresh token is invalid")
	}

	claims, err := j.ExtractClaims(token, tokenSecret)
	if err != nil {
		return "", err
	}

	return j.SignToken(claims, tokenSecret, expSec)
}

func NewJwtService() JwtService {
	return &jwtService{}
}

func (*jwtService) SignToken(mapClaims MapClaims, secret string, expSec time.Duration) (string, error) {
	c := jwt.MapClaims(mapClaims)
	c["exp"] = expSec.Seconds()
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tk.SignedString([]byte(secret))
}

func (j *jwtService) IsValid(t string, secret string) bool {
	tk, err := j.extractToken(t, secret)
	if err != nil {
		return false
	}

	if _, ok := tk.Claims.(jwt.Claims); !ok && !tk.Valid {
		fmt.Println("token not valid")
		return false
	}
	return true
}

func (j *jwtService) ExtractClaims(t string, secret string) (MapClaims, error) {
	claims, err := j.extractToken(t, secret)
	if err != nil {
		return nil, err
	}

	c := claims.Claims.(jwt.MapClaims)
	return MapClaims(c), nil
}

func (*jwtService) extractToken(t string, secret string) (*jwt.Token, error) {
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method, %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}
