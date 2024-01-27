package request

type CreateUserRequest struct {
	Name     string `validate:"required,min=1" json:"name" binding:"required"`
	Email    string `validate:"required,min=1,email" json:"email" binding:"required"`
	Password string `json:"password" validate:"required,min=8"`
}
