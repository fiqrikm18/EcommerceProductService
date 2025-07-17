package usecase

import (
	"ecommerce/internal/domain/brand/dto"
	"ecommerce/internal/domain/brand/entity"
	"ecommerce/internal/domain/brand/repository"
	"errors"
	"math"
)

type IBrandUseCase interface {
	FindAll(params *dto.BrandPaginationDTO) (int, int, []*dto.FindBrandDTO, error)
	FindById(payload *dto.BrandWithIdDTO) (*dto.FindBrandDTO, error)
	CreateBrand(payload *dto.CreateBrandDTO) error
	UpdateBrand(payload *dto.UpdateBrandDTO) error
	DeleteBrand(payload *dto.BrandWithIdDTO) error
}

type BrandUseCase struct {
	repository repository.IBrandRepository
}

func NewBrandUseCase(repository repository.IBrandRepository) *BrandUseCase {
	return &BrandUseCase{
		repository: repository,
	}
}

func (uc *BrandUseCase) CreateBrand(payload *dto.CreateBrandDTO) error {
	brand := &entity.Brand{
		Name: payload.Name,
	}

	err := uc.repository.Create(brand)
	if err != nil {
		return err
	}

	return nil
}

func (uc *BrandUseCase) UpdateBrand(payload *dto.UpdateBrandDTO) error {
	brand := &entity.Brand{
		Name: payload.Name,
		ID:   uint(payload.ID),
	}

	brandExist, err := uc.repository.FindById(brand.ID)
	if err != nil {
		return err
	}

	if brandExist == nil {
		return errors.New("brand not found")
	}

	err = uc.repository.Update(brand)
	if err != nil {
		return err
	}

	return nil
}

func (uc *BrandUseCase) DeleteBrand(payload *dto.BrandWithIdDTO) error {
	brand := &entity.Brand{
		ID: uint(payload.ID),
	}

	brandExist, err := uc.repository.FindById(brand.ID)
	if err != nil {
		return err
	}

	if brandExist == nil {
		return errors.New("brand not found")
	}

	err = uc.repository.Delete(brand)
	if err != nil {
		return err
	}

	return nil
}

func (uc *BrandUseCase) FindById(payload *dto.BrandWithIdDTO) (*dto.FindBrandDTO, error) {
	brand, err := uc.repository.FindById(uint(payload.ID))
	if err != nil {
		return nil, err
	}

	if brand == nil {
		return nil, errors.New("brand not found")
	}

	return &dto.FindBrandDTO{
		ID:        int64(brand.ID),
		Name:      brand.Name,
		CreatedAt: brand.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: brand.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (uc *BrandUseCase) FindAll(params *dto.BrandPaginationDTO) (int, int, []*dto.FindBrandDTO, error) {
	brandsDto := make([]*dto.FindBrandDTO, 0)

	if params.Page == 0 {
		params.Page = 1
	}

	if params.PerPage == 0 {
		params.PerPage = 10
	}

	if params.Sort == "" {
		params.Sort = "desc"
	}

	if params.SortBy == "" {
		params.SortBy = "created_at"
	}

	brands, err := uc.repository.FindAll(params)
	if err != nil {
		return 0, 0, make([]*dto.FindBrandDTO, 0), err
	}

	if len(brands) > 0 {
		for _, b := range brands {
			brandsDto = append(brandsDto, &dto.FindBrandDTO{
				ID:        int64(b.ID),
				Name:      b.Name,
				CreatedAt: b.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: b.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
	}

	totalPage := 0.0
	count, err := uc.repository.Count()
	if err != nil {
		return 0, 0, brandsDto, err
	}

	totalPage = math.Ceil(float64(count) / float64(params.PerPage))
	return int(count), int(totalPage), brandsDto, nil
}
