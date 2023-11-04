package auth

func validateSessionCookie(userID int, value string) bool {
	expectedValue, exists := sessionStore[userID]
	if !exists {
		return false
	}
	return value == expectedValue
}
