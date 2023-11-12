package structs

type User struct {
	Id       int
	Nama     string
	Username string
	Email    string
	Password string
	Auth     string
}

type Role struct {
	Id          int
	Role        string
	Description string
}

type Previlage struct {
	Id          int
	Previlage   string
	Description string
}
