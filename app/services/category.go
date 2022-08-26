package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-sample/app/models"
	"go-sample/app/repositories"
	"go-sample/app/utils/logic"
	"go-sample/app/utils/validation"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type CategoryService interface {
	GetAllCategories(w http.ResponseWriter, userId int) ([]models.BaseCategoryResponse, error)
	GetCategoryById(w http.ResponseWriter, r *http.Request, userId int) (models.BaseCategoryResponse, error)
	CreateCategory(w http.ResponseWriter, r *http.Request, userId int) (models.BaseCategoryResponse, error)
	DeleteCategory(w http.ResponseWriter, r *http.Request, userId int) error
	UpdateCategory(w http.ResponseWriter, r *http.Request, userId int) (models.BaseCategoryResponse, error)
	SendAllCategoryResponse(w http.ResponseWriter, categories *[]models.BaseCategoryResponse)
	SendCategoryResponse(w http.ResponseWriter, todo *models.BaseCategoryResponse)
	SendCreateCategoryResponse(w http.ResponseWriter, todo *models.BaseCategoryResponse)
	SendDeleteCategoryResponse(w http.ResponseWriter)
}

type categoryService struct {
	cr repositories.CategoryRepository
	cl logic.CategoryLogic
	rl logic.ResponseLogic
	cv validation.CategoryValidation
}

func NewCategoryService(cr repositories.CategoryRepository, cl logic.CategoryLogic, rl logic.ResponseLogic, cv validation.CategoryValidation) CategoryService {
	return &categoryService{cr, cl, rl, cv}
}

/*
 Categoryリストを取得
*/
func (ts *categoryService) GetAllCategories(w http.ResponseWriter, userId int) ([]models.BaseCategoryResponse, error) {
	var categories []models.Category

	// categoryリストデータ取得
	if err := ts.cr.GetAllCategories(&categories, userId); err != nil {
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("データ取得に失敗"), http.StatusInternalServerError)
		return nil, err
	}
	// レスポンス用の構造体に変換
	responseCategories := ts.cl.CreateAllCategoryResponse(&categories)

	return responseCategories, nil
}

/*
 IDに紐づくCategoryを取得
*/
func (ts *categoryService) GetCategoryById(w http.ResponseWriter, r *http.Request, userId int) (models.BaseCategoryResponse, error) {
	// getパラメータからIDを取得
	vars := mux.Vars(r)
	id := vars["id"]
	var category models.Category
	// Categoryデータ取得処理
	if err := ts.cr.GetCategoryById(&category, id); err != nil {
		var errMessage string
		var statusCode int
		// https://gorm.io/ja_JP/docs/error_handling.html
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusBadRequest
			errMessage = "該当データは存在しません。"
		} else {
			statusCode = http.StatusInternalServerError
			errMessage = "データ取得に失敗しました。"
		}
		// エラーレスポンス送信
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseCategoryResponse{}, err
	}

	// レスポンス用の構造体に変換
	responseTodos := ts.cl.CreateCategoryResponse(&category)

	return responseTodos, nil
}

/*
 Category新規登録処理
*/
func (cs *categoryService) CreateCategory(w http.ResponseWriter, r *http.Request, userId int) (models.BaseCategoryResponse, error) {
	// ioutil: ioに特化したパッケージ
	reqBody, _ := ioutil.ReadAll(r.Body)
	var mutationCategoryRequest models.MutationCategoryRequest
	if err := json.Unmarshal(reqBody, &mutationCategoryRequest); err != nil {
		log.Fatal(err)
		errMessage := "リクエストパラメータを構造体へ変換処理でエラー発生"
		cs.rl.SendResponse(w, cs.rl.CreateErrorStringResponse(errMessage), http.StatusInternalServerError)
		return models.BaseCategoryResponse{}, err
	}
	// バリデーション
	if err := cs.cv.MutationCategoryValidate(mutationCategoryRequest); err != nil {
		// バリデーションエラーのレスポンスを送信
		cs.rl.SendResponse(w, cs.rl.CreateErrorResponse(err), http.StatusBadRequest)
		return models.BaseCategoryResponse{}, err
	}

	var category models.Category
	category.Title = mutationCategoryRequest.Title

	// categoryデータ新規登録処理
	if err := cs.cr.CreateCategory(&category); err != nil {
		cs.rl.SendResponse(w, cs.rl.CreateErrorStringResponse("データ取得に失敗しました。"), http.StatusInternalServerError)
		return models.BaseCategoryResponse{}, err
	}

	// 登録したcategoryデータ取得処理
	if err := cs.cr.GetLastCategory(&category); err != nil {
		var errMessage string
		var statusCode int
		// https://gorm.io/ja_JP/docs/error_handling.html
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusBadRequest
			errMessage = "該当データは存在しません。"
		} else {
			statusCode = http.StatusInternalServerError
			errMessage = "データ取得に失敗しました。"
		}
		// エラーレスポンス送信
		cs.rl.SendResponse(w, cs.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseCategoryResponse{}, err
	}

	// レスポンス用の構造体に変換
	responseCategories := cs.cl.CreateCategoryResponse(&category)

	return responseCategories, nil
}

/*
 Category削除処理
*/
func (ts *categoryService) DeleteCategory(w http.ResponseWriter, r *http.Request, userId int) error {
	// getパラメータからIDを取得
	vars := mux.Vars(r)
	id := vars["id"]
	var category models.Category
	// データ削除処理
	if err := ts.cr.DeleteCategory(id, &category); err != nil {
		fmt.Println(err)
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponseCategory(err), http.StatusInternalServerError)
		return err
	}
	return nil
}

/*
 Category更新処理
*/
func (ts *categoryService) UpdateCategory(w http.ResponseWriter, r *http.Request, userId int) (models.BaseCategoryResponse, error) {
	// GetパラメータからIDを取得
	vars := mux.Vars(r)
	id := vars["id"]
	// request bodyから値を取得
	reqBody, _ := ioutil.ReadAll(r.Body)

	var mutationCategoryRequest models.MutationCategoryRequest
	if err := json.Unmarshal(reqBody, &mutationCategoryRequest); err != nil {
		fmt.Print("======")
		log.Fatal(err)
		errMessage := "リクエストパラメータを構造体へ変換処理でエラー発生"
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse(errMessage), http.StatusInternalServerError)
		return models.BaseCategoryResponse{}, err
	}
	// バリデーション
	if err := ts.cv.MutationCategoryValidate(mutationCategoryRequest); err != nil {
		// バリデーションエラーのレスポンスを送信
		ts.rl.SendResponse(w, ts.rl.CreateErrorResponse(err), http.StatusBadRequest)
		return models.BaseCategoryResponse{}, err
	}

	// 更新用データ用意
	var updateCategory models.Category
	updateCategory.Title = mutationCategoryRequest.Title

	// Categoryデータ新規登録処理
	if err := ts.cr.UpdateCategory(&updateCategory, id); err != nil {
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse("データ更新に失敗しました。"), http.StatusInternalServerError)
		return models.BaseCategoryResponse{}, err
	}

	// // 更新データを取得
	var category models.Category
	if err := ts.cr.GetCategoryById(&category, id); err != nil {
		var errMessage string
		var statusCode int
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusBadRequest
			errMessage = "該当データは存在しません。"
		} else {
			statusCode = http.StatusInternalServerError
			errMessage = "データ取得に失敗しました。"
		}
		// エラーレスポンス送信
		ts.rl.SendResponse(w, ts.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseCategoryResponse{}, err
	}

	// レスポンス用の構造体に変換
	responseCategories := ts.cl.CreateCategoryResponse(&updateCategory)

	return responseCategories, nil
}

/*
 Categoryリスト取得APIのレスポンス送信処理
*/
func (ts *categoryService) SendAllCategoryResponse(w http.ResponseWriter, todos *[]models.BaseCategoryResponse) {
	var response models.AllCategoryResponse
	response.Categories = *todos
	// レスポンスデータ作成
	responseBody, _ := json.Marshal(response)

	// レスポンス送信
	ts.rl.SendResponse(w, responseBody, http.StatusOK)
}

/*
 Categoryデータのレスポンス送信処理
*/
func (ts *categoryService) SendCategoryResponse(w http.ResponseWriter, todo *models.BaseCategoryResponse) {
	var response models.CategoryResponse
	response.Category = *todo
	// レスポンスデータ作成
	responseBody, _ := json.Marshal(response)
	// レスポンス送信
	ts.rl.SendResponse(w, responseBody, http.StatusOK)
}

/*
 CreateCategoryAPIのレスポンス送信処理
*/
func (ts *categoryService) SendCreateCategoryResponse(w http.ResponseWriter, category *models.BaseCategoryResponse) {
	var response models.CategoryResponse
	response.Category = *category
	// レスポンスデータ作成
	responseBody, _ := json.Marshal(response)
	// レスポンス送信
	ts.rl.SendResponse(w, responseBody, http.StatusCreated)
}

/*
 DeleteCategoryAPIのレスポンス送信処理
*/
func (ts *categoryService) SendDeleteCategoryResponse(w http.ResponseWriter) {
	// レスポンス送信
	ts.rl.SendNotBodyResponse(w)
}
