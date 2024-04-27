package service

import (
	"github.com/JLavrin/project-aigo/db"
	"github.com/JLavrin/project-aigo/src/domain/article/model"
)

func Get(id int) (*model.Article, error) {
	article := &model.Article{}
	err := db.DB.Where("id = ?", id).First(article).Error

	return article, err
}

func GetAll() ([]model.Article, error) {
	var articles []model.Article
	err := db.DB.Find(&articles).Error

	return articles, err
}

func Create(a *model.Article) (int, error) {
	err := db.DB.Create(a).Error

	return int(a.ID), err
}

func Update(a *model.Article, id *int) error {
	err := db.DB.Model(&model.Article{}).Where("id = ?", id).Updates(a).Error

	return err
}

func Delete(id *int) error {
	err := db.DB.Delete(&model.Article{}, id).Error

	return err
}
