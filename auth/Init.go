package auth

const (
	sessionCookieName = "Session"
)

var (
	sessionStore map[string]int
)

func init() {
	sessionStore = make(map[string]int)
}
