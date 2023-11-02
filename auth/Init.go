package auth

const (
	idCookieName      = "UserID"
	sessionCookieName = "Session"
)

var (
	sessionStore map[int]string
)

func init() {
	sessionStore = make(map[int]string)
}
