package models

type LogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// {
//   "username":"grey123",
//   "email":"grey@123",
//   "password":"12345"
// }
