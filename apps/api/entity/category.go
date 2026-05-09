package entity

import "time"

type Category struct {
	ID            string    `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	IsSystem      bool      `json:"is_system" db:"is_system"`
	IsArchived    bool      `json:"is_archived" db:"is_archived"`
	MarkForDelete bool      `json:"mark_for_delete" db:"mark_for_delete"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type SubCategory struct {
	ID            string    `json:"id" db:"id"`
	CategoryID    string    `json:"category_id" db:"category_id"`
	UserID        string    `json:"user_id" db:"user_id"`
	Name          string    `json:"name" db:"name"`
	IsSystem      bool      `json:"is_system" db:"is_system"`
	IsArchived    bool      `json:"is_archived" db:"is_archived"`
	MarkForDelete bool      `json:"mark_for_delete" db:"mark_for_delete"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type CategoryWithSubCategory struct {
	CategoryID      string `json:"category_id" db:"category_id"`
	SubCategoryID   string `json:"subcategory_id" db:"subcategory_id"`
	CategoryName    string `json:"category_name" db:"category_name"`
	SubCategoryName string `json:"subcategory_name" db:"subcategory_name"`
}
