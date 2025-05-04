package service

import (
	"context"

	"github.com/byvinesse/vinance-backend/entity"
	"github.com/byvinesse/vinance-backend/model"
	"github.com/byvinesse/vinance-backend/repository"
)

type CategoryService struct {
	categoryRepo repository.Category
}

func NewCategoryService(categoryRepo repository.Category) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *CategoryService) GetCompleteCategory(ctx context.Context, userID string) ([]model.GetCompleteCategoriesResponse, error) {
	res, err := s.categoryRepo.FindCompleteCategory(ctx, userID)
	if err != nil {
		return nil, err
	}

	return toGetCompleteCategoryResponse(res), nil
}

func toGetCompleteCategoryResponse(responses []entity.CategoryWithSubCategory) []model.GetCompleteCategoriesResponse {
	var res []model.GetCompleteCategoriesResponse

	for _, response := range responses {
		res = append(res, model.GetCompleteCategoriesResponse{
			CategoryID:      response.CategoryID,
			SubCategoryID:   response.SubCategoryID,
			CategoryName:    response.CategoryName,
			SubCategoryName: response.SubCategoryName,
		})
	}

	return res
}
