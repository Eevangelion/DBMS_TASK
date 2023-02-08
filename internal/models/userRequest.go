package models

type UserRequest struct {
	Name                string `json:"name"`
	Email               string `json:"email"`
	TransformedPassword string `json:"transformed_password"`
}
