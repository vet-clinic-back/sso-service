package models

type User struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullname,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"` // password hash
}

type Owner struct {
	User
}

type Vet struct {
	User
	Position     string `json:"position,omitempty"`
	ClinicNumber string `json:"clinic_number,omitempty"`
}
