package models

type Todo struct {
	BaseModel
	Title      string `gorm:"size:255" json:"title,omitempty"`
	Comment    string `gorm:"type:text" json:"comment,omitempty"`
	UserId     int    `gorm:"not null" json:"user_id"`
	User       User   `gorm:"foreignKey:UserId"`
	CategoryId int    `gorm:"not null" json:"category_id"`
}

type MutationTodoRequest struct {
	Title      string `json:"title,omitempty"`
	Comment    string `json:"comment,omitempty"`
	CategoryId int    `json:"category_id,omitempty"`
}

type BaseTodoResponse struct {
	BaseModel
	Title      string `gorm:"size:255" json:"title,omitempty"`
	Comment    string `gorm:"type:text" json:"comment,omitempty"`
	CategoryId int    `gorm:"type:int" json:"category_id"`
}

type TodoResponse struct {
	Todo BaseTodoResponse `json:"todo"`
}

type AllTodoResponse struct {
	Todos []BaseTodoResponse `json:"todos"`
}
