package token

import (
	"crypto/rand"
	"io"
)

func GenerateOTP(length int) string {
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

//const otpChars = "1234567890"
//
//func GenerateOTP(length int) (string, error) {
//    buffer := make([]byte, length)
//    _, err := rand.Read(buffer)
//    if err != nil {
//        return "", err
//    }
//
//    otpCharsLength := len(otpChars)
//    for i := 0; i < length; i++ {
//        buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
//    }
//
//    return string(buffer), nil
//}
