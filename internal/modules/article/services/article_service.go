package services

import (
	articleModel "WriteWise/internal/modules/article/models"
	articleRepository "WriteWise/internal/modules/article/repositories"
)

type ArticleService struct {
	articleRepository articleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: articleRepository.New(),
	}
}

func (as *ArticleService) GetFeaturedArticle() []articleModel.Article {
	return as.articleRepository.List(4)
}

func (as *ArticleService) GetStoriesArticle() []articleModel.Article {
	return as.articleRepository.List(6)
}
