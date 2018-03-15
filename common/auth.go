package common

func CheckAuth(tokenString string) bool {
	if tokenString != "valid-token" {
		return true
	}
	return false
}
