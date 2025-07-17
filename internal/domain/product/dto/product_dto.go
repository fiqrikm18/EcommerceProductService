package dto

import "ecommerce/internal/domain/brand/dto"

type CreateProductDTO struct {
	Name    string `json:"name" validate:"required"`
	Price   int    `json:"price" validate:"required,numeric"`
	Qty     int    `json:"qty" validate:"required,numeric"`
	BrandId int64  `json:"brand_id" validate:"required,numeric"`
}

type UpdateProductDTO struct {
	ID      int64  `json:"id" swaggerignore:"true"`
	Name    string `json:"name"`
	Price   int    `json:"price" validate:"numeric"`
	Qty     int    `json:"qty" validate:"numeric"`
	BrandId int64  `json:"brand_id" validate:"numeric"`
}

type ProductWithIdDTO struct {
	ID int64 `json:"id" form:"id" param:"id" query:"id"`
}

type FindProductDTO struct {
	ID        int64             `json:"id"`
	Name      string            `json:"name"`
	Price     int               `json:"price"`
	Qty       int               `json:"qty"`
	Brand     *dto.FindBrandDTO `json:"brand"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}

type ProductPaginationDTO struct {
	PerPage int64  `json:"per_page" query:"per_page" validate:"required,number"`
	Page    int64  `json:"page" query:"page" validate:"required,number"`
	Sort    string `json:"sort" query:"sort" validate:"required,oneof=asc desc"`
	SortBy  string `json:"sort_by" query:"sort_by"`
	Search  string `json:"search" query:"search"`
}
