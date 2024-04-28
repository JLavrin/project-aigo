package model

type Article struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Title         string `json:"title" binding:"required" gorm:"type:varchar(100)"`
	Content       string `json:"content" binding:"required" gorm:"type:text"`
	CoverImageUrl string `json:"coverImageUrl" binding:"required" gorm:"type:varchar(255)"`
}

type Articles []Article
