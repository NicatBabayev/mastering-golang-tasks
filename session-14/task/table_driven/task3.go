package table_driven

import (
	"fmt"
	"regexp"
)

func ReverseString(s string) (string, error) {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	res := make([]byte, 0)
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	if len(s) == 0 {
		return string(res), fmt.Errorf("Empty string detected")
	}
	if re.MatchString(s) {
		return string(res), fmt.Errorf("Special character detected")
	}
	if string(res) == s {
		return string(res), fmt.Errorf("Palindrome detected.")
	}

	return string(res), nil
}

func RunTask3() {

	fmt.Println(ReverseString("D@@R"))

}
