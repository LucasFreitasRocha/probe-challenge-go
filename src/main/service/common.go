package service

import "regexp"

func validateString(s string, r string) bool {
	re := regexp.MustCompile(r)
	return re.MatchString(s)
}