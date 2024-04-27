package article_service

import "github.com/JLavrin/project-aigo/models"

type Article struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
}

func (a *Article) Create() error {
	article := map[string]interface{}{
		"title":         a.Title,
		"content":       a.Content,
		"coverImageUrl": a.CoverImageUrl,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}

	return nil
}

func (a *Article) Update() bool {
	return true
}

func (a *Article) Get() *Article {
	return nil
}

func (a *Article) GetAll() []Article {

	return []Article{}
}

func (a *Article) Delete() bool {
	return true
}
