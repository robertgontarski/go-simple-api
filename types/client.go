package types

type CreateClientRequest struct {
	Email    string `json:"email" validate:"required,email" err_required_msg:"field is required" err_email_msg:"invalid email address"`
	Password string `json:"password" validate:"required" err_required_msg:"field is required"`
}
