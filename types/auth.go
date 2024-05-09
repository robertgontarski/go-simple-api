package types

type LoginAuthRequest struct {
	Email    string `json:"email" validate:"required" err_required_msg:"field is required"`
	Password string `json:"password" validate:"required" err_required_msg:"field is required"`
}
