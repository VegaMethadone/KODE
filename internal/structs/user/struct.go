package user

type User struct {
	Login    string `json:"login"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
