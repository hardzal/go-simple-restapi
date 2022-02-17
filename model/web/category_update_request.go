package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required,min=3,max=150" json:"name"`
}
