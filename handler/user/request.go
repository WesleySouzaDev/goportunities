package user

type CreateUserRequest struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func (r *CreateUserRequest) Validate() error {

	return nil
}
