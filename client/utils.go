package client

import (
	"regexp"
)

func isPwdLegal(password string) bool {
	match, _ := regexp.MatchString("^[0-9]{6,12}$", password)
	return match
}
