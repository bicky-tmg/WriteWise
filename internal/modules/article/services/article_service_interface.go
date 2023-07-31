package services

import ArticleResponse "WriteWise/internal/modules/article/responses"

type ArticleServiceInterface interface {
	GetFeaturedArticle() ArticleResponse.Articles
	GetStoriesArticle() ArticleResponse.Articles
}
