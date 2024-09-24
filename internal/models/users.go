package models

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Phone      string `json:"phone,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"` // password hash
	Role       string `json:"role,omitempty"`     // TODO enum may be?
	Hospital   string `json:"hospital,omitempty"`
}
