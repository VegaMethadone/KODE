package token_test

import (
	"errors"
	"fmt"
	"kode/internal/structs/user"
	"kode/internal/token"
	"testing"
	"time"
)

var testCase = &user.User{
	Login:    "Stany",
	Mail:     "Ivanov@gmail.com",
	Password: "qweqwe",
}

func TestJWTFunctions(t *testing.T) {
	tokenString, err := token.NewJwtToken(1, testCase.Login, testCase.Mail, time.Now(), time.Now().Add(time.Hour*24))
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(tokenString)

	claims, err := token.ExtractJWT(tokenString)
	if err != nil {
		t.Fatal(err)
		return
	}

	if testCase.Login != claims["Login"] || testCase.Mail != claims["Mail"] {
		t.Fatal(errors.New("login or mail is not the same"))
		fmt.Printf("%s\t%s\n", testCase.Login, claims["Login"])
		fmt.Printf("%s\t%s\n", testCase.Mail, claims["Mail"])
		return
	}
	fmt.Printf("%s\n%s\n", claims["Login"], claims["Mail"])
}
