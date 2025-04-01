package utils

import "regexp"

func validateNumbers(input string) bool {
	regex := "^[0-9]+$"
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}

func validateCode39(input string) bool {
	regex := "^[A-Z0-9\\-\\.\\$\\/\\+\\%]+$"
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}

func validateCode93(input string) bool {
	regex := "^[A-Z0-9!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~]+$"
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}

func validateASCII(input string) bool {
	regex := "^[\\x00-\\x7F]+$"
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}

func validateAztec(input string) bool {
	// (A-Z, a-z), (0-9)
	// space, +, -, ., $, /, :, ;, etc.
	regex := "^[A-Za-z0-9\\-\\.\\$\\/\\:\\+\\,\\?\\!\\*\\(\\)]+$"
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}

func validateCodabar(input string) bool {
	// Codabar: start and end [A, B, C, D], in the mid nums and sym
	regex := "^[ABCD][0-9\\-\\.\\$\\/\\:\\+]+[ABCD]$"
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}
