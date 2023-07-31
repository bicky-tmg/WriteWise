package repositories

import (
	articleModel "WriteWise/internal/modules/article/models"
)

type ArticleRepositoryInterface interface {
	List(limit int) []articleModel.Article
}
