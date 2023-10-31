package category_repository

import (
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
)

type CategoryRepository interface {
	CreateCategory(categoryPayLoad *entity.Category) (*dto.CreateNewCategoriesResponse, errs.Error)
	GetCategory() ([]CategoryProductMapped, errs.Error)
}
