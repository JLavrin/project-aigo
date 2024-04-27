package models

import "fmt"

type Article struct {
	Model

	Title         string `json:"title"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
}

func AddArticle(data map[string]interface{}) error {
	article := Article{
		Title:         data["title"].(string),
		Content:       data["content"].(string),
		CoverImageUrl: data["coverImageUrl"].(string),
	}

	fmt.Printf("%+v\n", article)

	if err := db.Create(&article).Error; err != nil {
		return err
	}

	return nil
}
