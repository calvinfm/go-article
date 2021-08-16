package model

import (
	"context"
	"go-article/entity"
	"go-article/settings"
)

type ArticleModel struct {
	db settings.DatabaseConfig
}

func (ArticleModel ArticleModel) AddArticle(article *entity.AddArticle) (bool, error) {
	sqlStatement := "SELECT * FROM add_article($1,$2,$3)"
	var isSuccess bool
	err := ArticleModel.db.GetDatabaseConfig().QueryRow(context.Background(), sqlStatement,
		article.Author,
		article.Title,
		article.Body,
	).Scan(
		&isSuccess,
	)

	return isSuccess, err
}

func (ArticleModel ArticleModel) GetArticle() ([]entity.Article, error) {
	sqlStatement := "SELECT * FROM get_article()"
	res, err := ArticleModel.db.GetDatabaseConfig().Query(context.Background(), sqlStatement)
	defer res.Close()

	if err != nil {
		return []entity.Article{}, err
	}

	articles := []entity.Article{}

	for res.Next() {
		article := entity.Article{}
		err2 := res.Scan(
			&article.Id,
			&article.Author,
			&article.Title,
			&article.Body,
			&article.Created,
		)

		if err2 != nil {
			return articles, err
		}

		articles = append(articles, article)
	}

	return articles, err
}
