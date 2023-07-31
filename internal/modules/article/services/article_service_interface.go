package services

import articleModel "WriteWise/internal/modules/article/models"

type ArticleServiceInterface interface {
	GetFeaturedArticle() []articleModel.Article
	GetStoriesArticle() []articleModel.Article
}
