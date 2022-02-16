package web

type CategoryCreateRequestRequest struct {
	Name string `validate:"required,min=3,max=150"`
}
