package repositories

import (
	articleModel "WriteWise/internal/modules/article/models"
	"WriteWise/pkg/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}

func (ar *ArticleRepository) List(limit int) []articleModel.Article {
	var articles []articleModel.Article

	ar.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)

	return articles
}

func (ar *ArticleRepository) Find(id int) articleModel.Article {
	var article articleModel.Article

	ar.DB.Joins("User").First(&article, id)

	return article
}
