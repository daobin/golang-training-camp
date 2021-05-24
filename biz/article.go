package biz

import (
	"errors"
	"trainingcamp/dao"
	"trainingcamp/pkg/e"
)

func GetArticles() ([]dao.Article, error) {
	articles, err := dao.GetArticles()
	if err != nil {
		if errors.Is(err, e.ErrNotFound) {
			// 查不到数据返回空，逻辑上不算错误
			return nil, nil
		} else {
			// 出现错误时，不关心，直接上抛给调用者处理
			return nil, err
		}
	}

	return articles, nil
}