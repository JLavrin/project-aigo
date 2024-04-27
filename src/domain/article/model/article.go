package model

type Article struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Title         string `json:"title" gorm:"type:varchar(100)"`
	Content       string `json:"content" gorm:"type:text"`
	CoverImageUrl string `json:"coverImageUrl" gorm:"type:varchar(255)"`
}

type Articles []Article
