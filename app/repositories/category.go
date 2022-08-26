package repositories

import (
	"go-sample/app/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories(category *[]models.Category, userId int) error
	GetCategoryById(category *models.Category, id string) error
	GetLastCategory(category *models.Category) error
	CreateCategory(category *models.Category) error
	DeleteCategory(id string, category *models.Category) error
	UpdateCategory(category *models.Category, id string) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

/*
  Categoryリストを取得
*/
func (tr *categoryRepository) GetAllCategories(category *[]models.Category, userId int) error {
	if err := tr.db.Find(&category).Error; err != nil {
		return err
	}

	return nil
}

/*
  Idに紐づくCategoryデータを取得
*/
func (tr *categoryRepository) GetCategoryById(category *models.Category, id string) error {
	if err := tr.db.Find(&category).Where("id=?", id).First(&category, id).Error; err != nil {
		return err
	}

	return nil
}

/*
 新規登録したCategoryデータを取得
*/
func (tr *categoryRepository) GetLastCategory(category *models.Category) error {
	if err := tr.db.Last(&category).Error; err != nil {
		return err
	}

	return nil
}

/*
 Category新規登録
*/
func (tr *categoryRepository) CreateCategory(category *models.Category) error {
	if err := tr.db.Create(&category).Error; err != nil {
		return err
	}

	return nil
}

/*
 Category削除処理
*/
func (tr *categoryRepository) DeleteCategory(id string, category *models.Category) error {

	db := tr.db.Where("id=?", id).Delete(&models.Category{})

	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected < 1 {

		return errors.Errorf("id=%w のCategoryデータが存在しません。", id)
	}

	return nil
}

/*
 Category更新処理
*/
func (tr *categoryRepository) UpdateCategory(category *models.Category, id string) error {
	if err := tr.db.Model(&category).Where("id=?", id).Updates(
		map[string]interface{}{
			"title": category.Title,
		}).Error; err != nil {
		return err
	}
	return nil
}
