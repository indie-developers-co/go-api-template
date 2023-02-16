package request

type User struct {
	Name     string `json:"name" validate:"required,alpha,min=2,max=100"`
	LastName string `json:"last_name" validate:"required,alpha,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
}
