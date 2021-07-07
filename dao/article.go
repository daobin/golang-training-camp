package dao

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"trainingcamp/pkg/e"
)

type Article struct {
	Dao

	Title string
	Content string
	State int
	CoverImageUrl string
	TagId int
	Tag Tag
}

func GetArticles() ([]Article, error){
	state := 1
	q := "select `id`, `tag_id`, `title`, `created_on`, `modified_on`, `deleted_on`, `cover_image_url`" +
		" from `blog_article` where `state` = ? limit 0, 100"

	rows, err := db.Query(q, state)
	if err != nil {
		return nil, errors.Wrap(e.ErrInternal, fmt.Sprintf("err: %s", err))
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		article := Article{}
		err := rows.Scan(&article.Id, &article.TagId, &article.Title, &article.CreatedOn, &article.ModifiedOn, &article.DeletedOn, &article.CoverImageUrl)
		if err != nil {
			return nil, errors.Wrap(e.ErrInternal, fmt.Sprintf("err: %s", err))
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func GetArticleById(id int) (*Article, error) {
	var article Article

	q := "select `id`, `tag_id`, `title`, `created_on`, `modified_on`, `deleted_on`, `cover_image_url`" +
		" from `blog_article` where `id` = ?"
	err := db.QueryRow(q, id).Scan(&article.Id, &article.TagId, &article.Title, &article.CreatedOn, &article.ModifiedOn, &article.DeletedOn, &article.CoverImageUrl)

	switch {
	case err == sql.ErrNoRows:
		return nil, errors.Wrap(e.ErrNotFound, fmt.Sprintf("err: %s", err))
	case err != nil:
		return nil, errors.Wrap(e.ErrInternal, fmt.Sprintf("err: %s", err))
	}

	return &article, nil
}
