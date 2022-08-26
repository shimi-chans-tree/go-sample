package validation

import (
	"go-sample/app/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CategoryValidation interface {
	MutationCategoryValidate(mutationCategoryRequest models.MutationCategoryRequest) error
}

type categoryValidation struct{}

func NewCategoryValidation() CategoryValidation {
	return &categoryValidation{}
}

/*
 Category新規登録、更新時のリクエストパラメータのバリデーション
*/

func (cv *categoryValidation) MutationCategoryValidate(mutationCategoryRequest models.MutationCategoryRequest) error {
	return validation.ValidateStruct(&mutationCategoryRequest,
		validation.Field(
			&mutationCategoryRequest.Title,
			validation.Required.Error("タイトルは必須入力です。"),
			validation.RuneLength(1, 10).Error("タイトルは 1～10 文字です"),
		),
	)
}
