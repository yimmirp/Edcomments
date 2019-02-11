package models

import "github.com/jinzhu/gorm"

//Comment es exportado
type Comment struct {
	gorm.Model
	UserID   uint      `json:"userId"`
	ParentID uint      `json:"username"`
	Vote     int32     `json:"votes"`
	Content  string    `json:"content"`
	HasVote  int8      `json:"hasVote" gorm:"-"`
	User     []User    `json:"user,omitempty"`
	Children []Comment `json:"children,omitempty"`
}
