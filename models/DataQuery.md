"Its Normal to get a method not allowed in the register endpoint"

query this  in the terminal to get the registered users: 
sqlite3 gorm.db "SELECT id, username, email, password FROM users;"

Try credentials:
{
  "username":"testuser",
 "email":"test@test.com",
 "password":"test1234"
}

{
  "message": "Registered successfully",
  "user": {
    "id": 7,
    "username": "testuser",
    "email": "test@test.com",
    "created_at": "2026-08-10T16:30:53.871558296+08:00"
  }
}


{
    "username": "aloha",
    "email": "@loha12",
    "password": "123"
    
}

{
    "message": "Registered successfully",
    "user": {
        "id": 10,
        "username": "aloha",
        "email": "@loha12",
        "created_at": "2026-08-11T07:16:48.584510223+08:00"
    }
}

{
  "email":"byerss",
  "password":"123"
}
----Separate your struct and your data----
Structs:
type RegisterRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
} 

Model/Data:
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"-"`
}