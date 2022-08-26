package models

import "time"

type Category struct {
	BaseModel
	Title string `gorm:"size:255" json:"title,omitempty"`
}

type BaseCategoryResponse struct {
	BaseModel
	Title string `gorm:"size:255" json:"title,omitempty"`
}

type MutationCategoryRequest struct {
	Title     string     `json:"title,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type CategoryResponse struct {
	Category BaseCategoryResponse `json:"category"`
}

type AllCategoryResponse struct {
	Categories []BaseCategoryResponse `json:"categories"`
}
