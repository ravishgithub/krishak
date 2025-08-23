package handlers

import "github.com/ravishgithub/krishak/authentication"

func ValidToken() string {
    token, _ := authentication.GenerateToken("admin")
    return token
}