package ws

type User struct {
	Code int      `json:"code"`
	Data UserData `json:"data"`
}

type Users struct {
	Code int        `json:"code"`
	Data []UserData `json:"data"`
}
type UserData struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}
