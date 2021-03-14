package encoder

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type HashEncoder interface {
	HashPass(pass string) (string, error)
	MatchHashPass(hashPass string, clearPass string) bool
}

type encoder struct{}

func (e encoder) HashPass(pass string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pass), 11)

	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(b), err
}

func (e encoder) MatchHashPass(hashPass string, clearPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(clearPass)) == nil
}

func NewHashEncoder() HashEncoder {
	return &encoder{}
}
