package todo

type User struct {
	id int `json:"-"`
	Name string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}