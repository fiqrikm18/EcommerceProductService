package usecase

import (
	BrandDto "ecommerce/internal/domain/brand/dto"
	BrandRepository "ecommerce/internal/domain/brand/repository"
	"ecommerce/internal/domain/product/dto"
	"ecommerce/internal/domain/product/entity"
	ProductRepository "ecommerce/internal/domain/product/repository"
	"errors"
	"math"
)

type IProductUseCase interface {
	FindAll(params *dto.ProductPaginationDTO) (int, int, []*dto.FindProductDTO, error)
	FindById(payload *dto.ProductWithIdDTO) (*dto.FindProductDTO, error)
	CreateProduct(product *dto.CreateProductDTO) error
	UpdateProduct(product *dto.UpdateProductDTO) error
	DeleteProduct(payload *dto.ProductWithIdDTO) error
}

type ProductUseCase struct {
	productRepository ProductRepository.IProductRepository
	brandRepository   BrandRepository.IBrandRepository
}

func NewProductUseCase(
	productRepository ProductRepository.IProductRepository,
	brandRepository BrandRepository.IBrandRepository,
) *ProductUseCase {
	return &ProductUseCase{
		productRepository: productRepository,
		brandRepository:   brandRepository,
	}
}

func (p *ProductUseCase) FindAll(params *dto.ProductPaginationDTO) (int, int, []*dto.FindProductDTO, error) {
	productDto := make([]*dto.FindProductDTO, 0)

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

	products, err := p.productRepository.FindAll(params)
	if err != nil {
		return 0, 0, make([]*dto.FindProductDTO, 0), err
	}

	if len(products) > 0 {
		for _, product := range products {
			brand, _ := p.brandRepository.FindById(uint(product.BrandId))
			productDto = append(productDto, &dto.FindProductDTO{
				ID:    int64(product.ID),
				Name:  product.Name,
				Price: product.Price,
				Qty:   product.Qty,
				Brand: &BrandDto.FindBrandDTO{
					ID:        int64(brand.ID),
					Name:      brand.Name,
					CreatedAt: brand.CreatedAt.Format("2006-01-02 15:04:05"),
					UpdatedAt: brand.UpdatedAt.Format("2006-01-02 15:04:05"),
				},
				CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
	}

	totalPage := 0.0
	count, err := p.productRepository.Count()
	if err != nil {
		return 0, 0, make([]*dto.FindProductDTO, 0), err
	}

	totalPage = math.Ceil(float64(count) / float64(params.PerPage))
	return count, int(totalPage), productDto, nil
}

func (p *ProductUseCase) FindById(payload *dto.ProductWithIdDTO) (*dto.FindProductDTO, error) {
	product, err := p.productRepository.FindById(int(payload.ID))
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("product not found")
	}

	brand, _ := p.brandRepository.FindById(uint(product.BrandId))
	return &dto.FindProductDTO{
		ID:    int64(product.ID),
		Name:  product.Name,
		Price: product.Price,
		Qty:   product.Qty,
		Brand: &BrandDto.FindBrandDTO{
			ID:        int64(brand.ID),
			Name:      brand.Name,
			CreatedAt: brand.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: brand.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
		CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (p *ProductUseCase) CreateProduct(payload *dto.CreateProductDTO) error {
	product := &entity.Product{
		Name:    payload.Name,
		Price:   payload.Price,
		Qty:     payload.Qty,
		BrandId: int(payload.BrandId),
	}

	err := p.productRepository.Create(product)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUseCase) UpdateProduct(payload *dto.UpdateProductDTO) error {
	product := &entity.Product{
		ID:      uint(payload.ID),
		Name:    payload.Name,
		Price:   payload.Price,
		Qty:     payload.Qty,
		BrandId: int(payload.BrandId),
	}

	productExists, err := p.productRepository.FindById(int(product.ID))
	if err != nil {
		return err
	}

	if productExists == nil {
		return errors.New("product not found")
	}

	err = p.productRepository.Update(product)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductUseCase) DeleteProduct(payload *dto.ProductWithIdDTO) error {
	product := &entity.Product{
		ID: uint(payload.ID),
	}

	productExists, err := p.productRepository.FindById(int(payload.ID))
	if err != nil {
		return err
	}

	if productExists == nil {
		return errors.New("product not found")
	}

	err = p.productRepository.Delete(product)
	if err != nil {
		return err
	}

	return nil
}
