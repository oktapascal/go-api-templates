package requests

type UserCreateRequest struct {
	IdNumber    string `validate:"required,min=1,max=16" json:"id_number"`
	Email       string `validate:"required,email" json:"email"`
	PhoneNumber string `validate:"required,min=10,max=13" json:"phone_number"`
	Address     string `validate:"required" json:"address"`
	FirstName   string `validate:"required,min=1,max=50" json:"first_name"`
	LastName    string `validate:"required,min=1,max=50" json:"last_name"`
	Provider    string `validate:"required" json:"provider"`
	ProviderId  int8   `validate:"required" json:"provider_id"`
}
