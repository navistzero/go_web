package server

import (
	// "fmt"
	"obedu/blogger/dao/db"
	"obedu/blogger/model"
)

// GetArticleRecordList 获取文章的简略信息和所属分类信息
func GetArticleRecordList(start, num int) (articleRecordList []*model.ArticleRecord, err error) {
	articleList, err := db.GetArticleList(start, num)
	if err != nil {
		return
	}
	categorys, err := db.GetAllCategory()
	if err != nil {
		return
	}
	categoryMap := make(map[int64]*model.Category)
	for _, category := range categorys {
		categoryMap[category.CategoryID] = category
	}
	// articleRecordList := make([]model.ArticleRecord, 0)
	for _, artcile := range articleList {
		var artcileRecord = &model.ArticleRecord{
			ArticleInfo: *artcile,
			Category:    *(categoryMap[artcile.ArticleCategory]),
		}
		articleRecordList = append(articleRecordList, artcileRecord)
	}
	return
}

// GetArticleRecordByCategory 获取某个分类下面的全部文章
func GetArticleRecordByCategory(categoryID int64,start, num int)(articleRecordList []*model.ArticleRecord, err error){
	articleList, err := db.GetArticleByCategory(categoryID,start, num)
	if err != nil {
		return
	}
	category, err := db.GetCategoryByID(categoryID)
	if err != nil {
		return
	}
	// articleRecordList := make([]model.ArticleRecord, 0)
	for _, artcile := range articleList {
		var artcileRecord = &model.ArticleRecord{
			ArticleInfo: *artcile,
			Category:    *category,
		}
		articleRecordList = append(articleRecordList, artcileRecord)
	}
	return 
}