package router

import (
	"go-sample/app/controllers"
	"go-sample/app/utils/logic"
	"net/http"

	"github.com/gorilla/mux"
)

type CategoryRouter interface {
	SetCategoryRouting(router *mux.Router)
}

type categoryRouter struct {
	tc controllers.CategoryController
}

func NewCategoryRouter(tc controllers.CategoryController) CategoryRouter {
	return &categoryRouter{tc}
}

func (tr *categoryRouter) SetCategoryRouting(router *mux.Router) {
	router.Handle("/api/v1/category", logic.JwtMiddleware.Handler(http.HandlerFunc(tr.tc.FetchAllCategories))).Methods("GET")
	router.Handle("/api/v1/category/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(tr.tc.FetchCategoryById))).Methods("GET")

	router.Handle("/api/v1/category", logic.JwtMiddleware.Handler(http.HandlerFunc(tr.tc.CreateCategory))).Methods("POST")
	router.Handle("/api/v1/category/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(tr.tc.DeleteCategory))).Methods("DELETE")
	router.Handle("/api/v1/category/{id}", logic.JwtMiddleware.Handler(http.HandlerFunc(tr.tc.UpdateCategory))).Methods("PUT")
}
