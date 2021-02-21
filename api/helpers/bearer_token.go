package helpers

func ExtractBerarerToken(auth string) string {
	return auth[7:]
}

func ValidateBearerToken(auth string) bool {
	if len(auth) != 39 {
		return false
	}
	return auth[:7] == "Bearer " && len(auth[7:]) == 32
}
