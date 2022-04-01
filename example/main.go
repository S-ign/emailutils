package main

import (
	"github.com/s-ign/emailutils/config"
	"github.com/s-ign/emailutils/email"
)

func main() {
	auth, err := config.CreateAuth(
		"test@test.com",
		"testtesttest",
	)
	if err != nil {
		return
	}
	message, err := email.NewMessage(
		"Why NeoVim is Awesome",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		"templates/template.html",
		[]string{"avery.carty7@gmail.com", "dave.carty@mojodomo.com"},
	)
	err = email.SendEmail(auth, config.OFFICE365, message)
	if err != nil {
		panic(err)
	}
}
