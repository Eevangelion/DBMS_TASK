package models

type UserRequestLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserRequestRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequestRegisterGithub struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Email               string `json:"email"`
	TransformedPassword string `json:"transformed_password"`
}
