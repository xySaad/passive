package twitter

import (
	"strings"
)

func IsValidUsername(username string) bool {
	if len(username) > 26 || len(username) == 0 {
		return false
	}
	for _, char := range strings.ToLower(username) {
		if char >= 'a' && char <= 'z' {
			continue
		}
		if char >= '0' && char <= '9' {
			continue
		}
		if char == '_' {
			continue
		}
		return false
	}
	return true
}
