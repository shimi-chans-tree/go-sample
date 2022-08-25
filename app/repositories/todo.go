package repositories

import (
	"fmt"
	"go-sample/app/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAllTodos(todo *[]models.Todo, userId int) error
	GetTodoById(todo *models.Todo, id string, userId int) error
	GetTodoLastByUserId(todo *models.Todo, userId int) error
	CreateTodo(todo *models.Todo) error
	DeleteTodo(id string, userId int) error
	UpdateTodo(todo *models.Todo, id string, userId int) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

/*
  Todoリストを取得
*/
func (tr *todoRepository) GetAllTodos(todo *[]models.Todo, userId int) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Find(&todo).Error; err != nil {
		return err
	}

	return nil
}

/*
  Idに紐づくTodoデータを取得
*/
func (tr *todoRepository) GetTodoById(todo *models.Todo, id string, userId int) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(&todo, id).Error; err != nil {
		return err
	}

	return nil
}

/*
 新規登録したTodoデータを取得
*/
func (tr *todoRepository) GetTodoLastByUserId(todo *models.Todo, userId int) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Last(&todo).Error; err != nil {
		return err
	}

	return nil
}

/*
 Todo新規登録
*/
func (tr *todoRepository) CreateTodo(todo *models.Todo) error {
	if err := tr.db.Create(&todo).Error; err != nil {
		return err
	}

	return nil
}

/*
 Todo削除処理
*/
func (tr *todoRepository) DeleteTodo(id string, userId int) error {

	db := tr.db.Where("id=?", id).Delete(&models.Todo{})
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected < 1 {

		fmt.Println(db.RowsAffected)
		return errors.Errorf("id=%w のTodoデータが存在しません。", id)
	}

	return nil
}

/*
 Todo更新処理
*/
func (tr *todoRepository) UpdateTodo(todo *models.Todo, id string, userId int) error {
	if err := tr.db.Model(&todo).Where("id=? AND user_id=?", id, userId).Updates(
		map[string]interface{}{
			"title":   todo.Title,
			"comment": todo.Comment,
		}).Error; err != nil {
		return err
	}
	return nil
}
