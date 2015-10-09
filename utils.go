package main

import "regexp"

func findMatch(reg, str string) string {
	re := regexp.MustCompile(reg)
	argF := re.FindString(str)
	if argF != "" {
		return argF
	}
	return ""
}
