package main

import (
	"strings"
)

func cleanInput(text string) []string {
	fmtString := strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
	return fmtString
}
