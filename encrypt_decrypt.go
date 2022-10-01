package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func maind() {
	// EncryptDecrypt("encrypt", "password")
	pass := "Psbest123"
	hashses := EncryptPassword(pass)
	fmt.Println(hashses)

	cek := CheckPassword(pass, hashses)
	fmt.Println(cek)

	cek = CheckPassword("dum", hashses)
	fmt.Println(cek)

}

func EncryptPassword(Origin string) string {
	var Hasil string
	var Err error
	_ = Err

	Hasil, Err = HashPassword(Origin) // ignore error for the sake of simplicity
	return Hasil
}

func CheckPassword(password string, hash string) bool {
	Hasil := CheckPasswordHash(password, hash)
	return Hasil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
