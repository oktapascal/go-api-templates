package responses

type UserResponse struct {
	IdNumber    string `json:"id_number"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	Provider    string `json:"provider"`
	ProviderId  int8   `json:"provider_id"`
}
