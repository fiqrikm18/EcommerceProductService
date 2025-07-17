package repository

import (
	"context"
	"ecommerce/config"
	"ecommerce/internal/domain/product/dto"
	"ecommerce/internal/domain/product/entity"
	"fmt"
	"log/slog"
)

//go:generate mockgen -source=product_repository.go -destination=mocks/product_repository_mock.go -package=mocks
type IProductRepository interface {
	Count() (int, error)
	FindAll(params *dto.ProductPaginationDTO) ([]*entity.Product, error)
	FindById(id int) (*entity.Product, error)
	Create(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(product *entity.Product) error
}

type ProductRepository struct {
	dbProvider *config.DatabaseConfiguration
	ctx        context.Context
	logger     *slog.Logger
}

func NewProductRepository(
	ctx context.Context,
	dbProvider *config.DatabaseConfiguration,
	logger *slog.Logger,
) *ProductRepository {
	return &ProductRepository{
		dbProvider: dbProvider,
		ctx:        ctx,
		logger:     logger,
	}
}

func (p *ProductRepository) Count() (int, error) {
	var count int64
	if err := p.dbProvider.WithContext(p.ctx).Model(&entity.Product{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (p *ProductRepository) FindAll(params *dto.ProductPaginationDTO) ([]*entity.Product, error) {
	products := make([]*entity.Product, 0)
	qw := p.dbProvider.WithContext(p.ctx).
		Model(&products).
		Limit(int(params.PerPage)).
		Offset(int(params.PerPage * (params.Page - 1))).
		Order(fmt.Sprintf("%s %s", params.SortBy, params.Sort))

	if err := qw.Find(&products).Error; err != nil {
		return make([]*entity.Product, 0), err
	}

	return products, nil
}

func (p *ProductRepository) FindById(id int) (*entity.Product, error) {
	product := &entity.Product{}
	if err := p.dbProvider.WithContext(p.ctx).Model(&entity.Product{}).Where("id = ?", id).First(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepository) Create(product *entity.Product) error {
	tx := p.dbProvider.WithContext(p.ctx).Begin()
	if err := tx.Create(product).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (p *ProductRepository) Update(product *entity.Product) error {
	tx := p.dbProvider.WithContext(p.ctx).Begin()
	if err := tx.Save(product).Where("id = ?", product.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (p *ProductRepository) Delete(product *entity.Product) error {
	tx := p.dbProvider.WithContext(p.ctx).Begin()
	if err := tx.Delete(product).Where("id = ?", product.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
