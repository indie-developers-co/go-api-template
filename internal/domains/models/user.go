package models

type User struct {
	Name     string `json:"name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Email    string `json:"email,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`
}
