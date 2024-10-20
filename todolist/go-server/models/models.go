package models

type ToDoList struct {
	ID     uint   `json:"id" gorm:"primaryKey"` 
	Task   string `json:"task,omitempty"`
	Status bool   `json:"status,omitempty"`
}
