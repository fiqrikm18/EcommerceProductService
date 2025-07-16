package repository

import (
	"context"
	"ecommerce/config"
	"ecommerce/internal/domain/brand/dto"
	"ecommerce/internal/domain/brand/entity"
	"fmt"
	"log/slog"
)

//go:generate mockgen -source=brand_repository.go -destination=mocks/brand_repository_mock.go -package=mocks
type IBrandRepository interface {
	Count() (int64, error)
	FindAll(params *dto.BrandPaginationDTO) ([]*entity.Brand, error)
	FindById(id uint) (*entity.Brand, error)
	Create(brand *entity.Brand) error
	Update(brand *entity.Brand) error
	Delete(brand *entity.Brand) error
}

type BrandRepository struct {
	dbProvider *config.DatabaseConfiguration
	ctx        context.Context
	logger     *slog.Logger
}

func NewBrandRepository(ctx context.Context, dbProvider *config.DatabaseConfiguration, logger *slog.Logger) *BrandRepository {
	return &BrandRepository{
		dbProvider: dbProvider,
		ctx:        ctx,
		logger:     logger,
	}
}

func (r *BrandRepository) Count() (int64, error) {
	var count int64
	if err := r.dbProvider.WithContext(r.ctx).Model(&entity.Brand{}).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (repo *BrandRepository) FindAll(params *dto.BrandPaginationDTO) ([]*entity.Brand, error) {
	brands := make([]*entity.Brand, 0)
	qw := repo.dbProvider.WithContext(repo.ctx).
		Model(&brands).
		Limit(int(params.PerPage)).
		Offset(int(params.PerPage * (params.Page - 1))).
		Order(fmt.Sprintf("%s %s", params.SortBy, params.Sort))

	if err := qw.Find(&brands).Error; err != nil {
		return make([]*entity.Brand, 0), err
	}

	return brands, nil
}

func (repo *BrandRepository) FindById(id uint) (*entity.Brand, error) {
	var brand *entity.Brand
	if err := repo.dbProvider.WithContext(repo.ctx).First(&brand, "id = ?", id).Error; err != nil {
		repo.logger.Error(err.Error())
		return nil, err
	}
	return brand, nil
}

func (repo *BrandRepository) Create(brand *entity.Brand) error {
	tx := repo.dbProvider.Begin()
	if err := tx.WithContext(repo.ctx).Create(brand).Error; err != nil {
		tx.Rollback()
		repo.logger.Error(err.Error())
		return err
	}
	tx.Commit()
	return nil
}

func (repo *BrandRepository) Update(brand *entity.Brand) error {
	tx := repo.dbProvider.Begin()
	if err := tx.WithContext(repo.ctx).Save(brand).Where("id = ?", brand.ID).Error; err != nil {
		tx.Rollback()
		repo.logger.Error(err.Error())
		return err
	}
	tx.Commit()
	return nil
}

func (repo *BrandRepository) Delete(brand *entity.Brand) error {
	tx := repo.dbProvider.Begin()
	if err := tx.WithContext(repo.ctx).Delete(brand).Error; err != nil {
		tx.Rollback()
		repo.logger.Error(err.Error())
		return err
	}
	tx.Commit()
	return nil
}
