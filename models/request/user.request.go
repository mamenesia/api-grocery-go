package request

type UserCreateRequest struct {
	Name    string `json:"name" validate:"required,min=3"`
	Email   string `json:"email" validate:"required,email,min=6,max=32"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
