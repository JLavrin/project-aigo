package model

type Article struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Title         string `json:"title" binding:"required" gorm:"type:varchar(100)"`
	Content       string `json:"content" gorm:"type:text"`
	CoverImageUrl string `json:"coverImageUrl" gorm:"type:varchar(255)"`
	IsPublished   bool   `json:"isPublished" gorm:"default:false"`
}

type Articles []Article
