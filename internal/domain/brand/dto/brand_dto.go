package dto

type CreateBrandDTO struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateBrandDTO struct {
	ID   int64  `json:"id" form:"id" param:"id" query:"id"`
	Name string `json:"name" form:"name"`
}

type BrandWithIdDTO struct {
	ID int64 `json:"id" form:"id" param:"id" query:"id"`
}

type FindBrandDTO struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BrandPaginationDTO struct {
	PerPage int64  `json:"per_page" query:"per_page" validate:"required,number"`
	Page    int64  `json:"page" query:"page" validate:"required,number"`
	Sort    string `json:"sort" query:"sort" validate:"required,oneof=asc desc"`
	SortBy  string `json:"sort_by" query:"sort_by"`
	Search  string `json:"search" query:"search"`
}
