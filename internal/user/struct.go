package user

// FormRegis as form data register
type FormRegis struct {
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
	RoleID   string `json:"role_id" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Status   string `json:"status"`
	Password string `json:"password" validate:"required"`
}

// FormLogin as form data register
type FormLogin struct {
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
