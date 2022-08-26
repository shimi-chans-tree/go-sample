package logic

import "go-sample/app/models"

type CategoryLogic interface {
	CreateAllCategoryResponse(categories *[]models.Category) []models.BaseCategoryResponse
	CreateCategoryResponse(categories *models.Category) models.BaseCategoryResponse
}

type categoryLogic struct{}

func NewCategoryLogic() CategoryLogic {
	return &categoryLogic{}
}

/*
 レスポンス用のTodoリストの構造体を作成
*/
func (cl *categoryLogic) CreateAllCategoryResponse(categories *[]models.Category) []models.BaseCategoryResponse {
	var responseCategories []models.BaseCategoryResponse
	for _, category := range *categories {
		var newCategory models.BaseCategoryResponse
		newCategory.BaseModel.ID = category.BaseModel.ID
		newCategory.BaseModel.CreatedAt = category.BaseModel.CreatedAt
		newCategory.BaseModel.UpdatedAt = category.BaseModel.UpdatedAt
		newCategory.BaseModel.DeletedAt = category.BaseModel.DeletedAt
		newCategory.Title = category.Title
		responseCategories = append(responseCategories, newCategory)
	}

	return responseCategories
}

/*
 レスポンス用のTodoの構造体を作成
*/
func (cl *categoryLogic) CreateCategoryResponse(category *models.Category) models.BaseCategoryResponse {
	var responseCategories models.BaseCategoryResponse
	responseCategories.BaseModel.ID = category.BaseModel.ID
	responseCategories.BaseModel.CreatedAt = category.BaseModel.CreatedAt
	responseCategories.BaseModel.UpdatedAt = category.BaseModel.UpdatedAt
	responseCategories.BaseModel.DeletedAt = category.BaseModel.DeletedAt
	responseCategories.Title = category.Title

	return responseCategories
}
