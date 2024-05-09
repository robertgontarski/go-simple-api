package types

type CreateProductRequest struct {
	Name  string  `json:"name" validate:"required" err_required_msg:"filed is required"`
	Price float64 `json:"price" validate:"gt=0" err_gt_msg:"value must be greater than 0"`
}
