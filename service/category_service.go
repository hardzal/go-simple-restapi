package service

import (
	"context"
	"restfulapi-golang/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequestRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}