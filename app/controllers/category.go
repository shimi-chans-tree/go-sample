package controllers

import (
	"go-sample/app/services"
	"net/http"
)

type CategoryController interface {
	FetchAllCategories(w http.ResponseWriter, r *http.Request)
	FetchCategoryById(w http.ResponseWriter, r *http.Request)
	CreateCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
}

type categoryController struct {
	ts services.TodoService
	as services.AuthService
	cs services.CategoryService
}

func NewCategoryController(ts services.TodoService, as services.AuthService, cs services.CategoryService) CategoryController {
	return &categoryController{ts, as, cs}
}

/*
 Categoryリスト取得
*/
func (tc *categoryController) FetchAllCategories(w http.ResponseWriter, r *http.Request) {
	// tokenからuserIdを所得
	userId, err := tc.as.GetUserIdFromToken(w, r)
	if userId == 0 || err != nil {
		return
	}

	// Categoryリスト取得処理
	alltodo, err := tc.cs.GetAllCategories(w, userId)
	if err != nil {
		return
	}

	// レスポンス送信処理
	tc.cs.SendAllCategoryResponse(w, &alltodo)
}

/*
 idに紐づくCategoryを取得
*/
func (tc *categoryController) FetchCategoryById(w http.ResponseWriter, r *http.Request) {
	// tokenからuserIdを所得
	userId, err := tc.as.GetUserIdFromToken(w, r)
	if userId == 0 || err != nil {
		return
	}
	// Categoryデータ取得処理
	responseTodo, err := tc.ts.GetTodoById(w, r, userId)
	if err != nil {
		return
	}

	// レスポンス送信処理
	tc.ts.SendTodoResponse(w, &responseTodo)
}

/*
 Category新規登録
*/
func (tc *categoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	// トークンからuserIdを取得
	userId, err := tc.as.GetUserIdFromToken(w, r)
	if userId == 0 || err != nil {
		return
	}

	// Categoryデータ取得処理
	responseCategory, err := tc.cs.CreateCategory(w, r, userId)
	if err != nil {
		return
	}

	// レスポンス送信処理
	tc.cs.SendCreateCategoryResponse(w, &responseCategory)
}

/*
 削除処理
*/
func (tc *categoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// トークンからuserIdを取得
	userId, err := tc.as.GetUserIdFromToken(w, r)
	if userId == 0 || err != nil {
		return
	}

	// データ削除処理
	if err := tc.cs.DeleteCategory(w, r, userId); err != nil {
		return
	}

	// レスポンス送信処理
	tc.ts.SendDeleteTodoResponse(w)
}

/*
 Category更新処理
*/
func (tc *categoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	// tokenからuserIdを所得
	userId, err := tc.as.GetUserIdFromToken(w, r)
	if userId == 0 || err != nil {
		return
	}

	// Category更新処理
	responseCategory, err := tc.cs.UpdateCategory(w, r, userId)
	if err != nil {
		return
	}

	// レスポンス送信処理
	tc.cs.SendCreateCategoryResponse(w, &responseCategory)
}
