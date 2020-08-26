package entity

import "time"

//Person struct
type Person struct {
	ID    uint64 `json:"id" gorm:"primary_key ; auto_increment"`
	Fname string `json:"first_name" gorm:"type:varchar(40);REQUIRED"`
	Lname string `json:"last_name" gorm:"type:varchar(40);REQUIRED"`
	Email string `json:"email" gorm:"type:varchar(40);UNIQUE;REQUIRED"`
}

// Video the struct of the video
type Video struct {
	ID          uint64    `json:"id" gorm:"primary_key ; auto_increment"`
	Title       string    `json:"title" binding:"min=5" validate:"is-cool" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:varchar(40)"`
	URL         string    `json:"url" gorm:"type:varchar(40) ; UNIQUE "`
	Author      Person    `json:"author" gorm:"type:varchar(40); foreign_key:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
