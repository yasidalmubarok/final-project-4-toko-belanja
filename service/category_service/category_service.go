package category_service

import (
	"net/http"
	"toko-belanja-app/dto"
	"toko-belanja-app/entity"
	"toko-belanja-app/pkg/errs"
	"toko-belanja-app/pkg/helpers"
	"toko-belanja-app/repository/category_repository"
)

type CategoryService interface {
	CreateCategory(categoryPayLoad *dto.CategoriesRequest) (*dto.CategoryResponse, errs.Error)
	GetAllCategory() (*dto.CategoryResponse, errs.Error)
	UpdateCategory(categoryId int, categoryPayLoad *dto.CategoriesRequest) (*dto.CategoryResponse, errs.Error)
	DeleteCategory(categoryId int) (*dto.CategoryResponse, errs.Error)
}

type categoryServiceImpl struct {
	cr category_repository.CategoryRepository
}

func NewCategoryService(categoryRepo category_repository.CategoryRepository) CategoryService {
	return &categoryServiceImpl{cr: categoryRepo}
}

func (cs *categoryServiceImpl) CreateCategory(categoryPayLoad *dto.CategoriesRequest) (*dto.CategoryResponse, errs.Error){
	err := helpers.ValidateStruct(categoryPayLoad)

	if err != nil {
		return nil, err
	}

	category := &entity.Category{
		Type: categoryPayLoad.Type,
	}

	response, err := cs.cr.CreateCategory(category)
	
	if err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		Code: http.StatusCreated,
		Message: "Category has been successfully created",
		Data: response,
	}, nil
}
func (cs *categoryServiceImpl) GetAllCategory() (*dto.CategoryResponse, errs.Error)
func (cs *categoryServiceImpl) UpdateCategory(categoryId int, categoryPayLoad *dto.CategoriesRequest) (*dto.CategoryResponse, errs.Error)
func (cs *categoryServiceImpl) DeleteCategory(categoryId int) (*dto.CategoryResponse, errs.Error)
