package services

import (
	ArticleRepository "WriteWise/internal/modules/article/repositories"
	ArticleResponse "WriteWise/internal/modules/article/responses"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (as *ArticleService) GetFeaturedArticle() ArticleResponse.Articles {
	articles := as.articleRepository.List(4)

	return ArticleResponse.ToArticles(articles)
}

func (as *ArticleService) GetStoriesArticle() ArticleResponse.Articles {
	articles := as.articleRepository.List(6)

	return ArticleResponse.ToArticles(articles)
}
