package token

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type MapClaims jwt.MapClaims

type JwtService interface {
	ExtractClaims(t, secret string) (MapClaims, error)
	SignToken(mapClaims *MapClaims, secret string, expSec time.Time) (string, error)
	IsValid(t, secret string) bool
	RefreshToken(
		refreshToken,
		refreshSecret,
		token,
		tokenSecret string,
		expSec time.Time,

	) (string, error)
	ValidateGRpc(
		ctx context.Context,
		refreshToken,
		refreshSecret,
		token,
		tokenSecret string,
		expSec time.Time,
	) error
}

type jwtService struct{}

func (j *jwtService) ValidateGRpc(
	ctx context.Context,
	refreshToken,
	refreshSecret,
	token,
	tokenSecret string,
	expSec time.Time,
) error {

	tk, err, expired := j.extractToken(token, tokenSecret)

	if expired {
		newT, er := j.RefreshToken(refreshToken, refreshSecret, token, tokenSecret, expSec)
		if er != nil {
			return er
		}

		return grpc.SetHeader(ctx, metadata.New(map[string]string{"u_auth": newT}))
	}

	if err != nil {
		return errors.New("auth token not valid")
	}

	if !tk.Valid {
		return errors.New("auth token not valid")
	}
	return nil
}

func (j *jwtService) RefreshToken(
	refreshToken,
	refreshSecret,
	token,
	tokenSecret string,
	exp time.Time,
) (string, error) {

	if ok := j.IsValid(refreshToken, refreshSecret); !ok {
		return "", errors.New("refresh token is invalid")
	}

	claims, err := j.ExtractClaims(token, tokenSecret)
	if err != nil {
		return "", err
	}

	return j.SignToken(&claims, tokenSecret, exp)
}

func (*jwtService) SignToken(mapClaims *MapClaims, secret string, exp time.Time) (string, error) {
	c := jwt.MapClaims(*mapClaims)
	c["exp"] = exp.Unix()
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tk.SignedString([]byte(secret))
}

func (j *jwtService) IsValid(t, secret string) bool {
	tk, err, _ := j.extractToken(t, secret)
	if err != nil {
		return false
	}

	if _, ok := tk.Claims.(jwt.Claims); !ok && !tk.Valid {
		fmt.Println("token not valid")
		return false
	}
	return true
}

func (j *jwtService) ExtractClaims(t, secret string) (MapClaims, error) {
	claims, err, _ := j.extractToken(t, secret)
	if err != nil {
		return nil, err
	}

	c := claims.Claims.(jwt.MapClaims)
	return MapClaims(c), nil
}

func (*jwtService) extractToken(t string, secret string) (*jwt.Token, error, bool) {
	tk, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		//if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		//	return nil, fmt.Errorf("unexpected signing method, %v", token.Header["alg"])
		//}
		return []byte(secret), nil
	})

	if err != nil && err.Error() == "Token is expired" {
		return tk, nil, true
	}

	return tk, err, false
}

func NewJwtService() JwtService {
	return &jwtService{}
}
