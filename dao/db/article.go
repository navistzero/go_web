package db

import (
	"fmt"
	"obedu/blogger/model"
)

// InsertArticle 插入文章
func InsertArticle(article *model.ArticleDetail) (articleID int64, err error) {
	sqlStr := "insert into article(category_id,title,summary,view_count,comment_count,user_name,content,create_time,updata_time) value(?,?,?,?,?,?,?,?,?)"
	stmt, err := DB.Prepare(sqlStr)
	if nil != err {
		fmt.Println("Prepare Failed!")
		return -1, err
	}
	// 将参数传递到SQL语句中执行
	res, err := stmt.Exec(article.ArticleCategory, article.ArticleTitle, article.ArticleSummary, article.ArticleViewCount, article.ArticleCommentCount, article.ArticleUserName, article.ArticleContent, article.ArticleCreateTime, article.ArticleUpdataTime)
	if nil != err {
		fmt.Println("Exec Failed!")
		return -1, err
	}
	articleID, err = res.LastInsertId()
	return
}

// GetArticleList huoqu
func GetArticleList(page, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if page >= pageSize || page < 0 {
		return
	}
	sqlStr := "select id,category_id,title,summary,view_count,comment_count,user_name,create_time,updata_time from article where status = 0 order by create_time desc limit ?,?"
	rows, err := DB.Query(sqlStr, page, pageSize)
	if err != nil {
		return
	}
	for rows.Next() {
		var myarticle model.ArticleInfo
		rows.Scan(&myarticle.ArticleID, &myarticle.ArticleCategory, &myarticle.ArticleTitle, &myarticle.ArticleSummary, &myarticle.ArticleViewCount, &myarticle.ArticleCommentCount, &myarticle.ArticleUserName, &myarticle.ArticleCreateTime, &myarticle.ArticleUpdataTime)
		articleList = append(articleList, &myarticle)
	}
	return
}

// GetArticleByID 单个
func GetArticleByID(id int64) (article *model.ArticleDetail, err error) {
	sqlStr := "select id,category_id,title,summary,view_count,comment_count,user_name,content,create_time,updata_time from article where id = ?"
	articleInfo := model.ArticleInfo{}
	article = &model.ArticleDetail{ArticleInfo: articleInfo}
	err = DB.QueryRow(sqlStr, id).Scan(&article.ArticleID, &article.ArticleCategory, &article.ArticleTitle, &article.ArticleSummary, &article.ArticleViewCount, &article.ArticleCommentCount, &article.ArticleUserName, &article.ArticleContent, &article.ArticleCreateTime, &article.ArticleUpdataTime)
	if err != nil {
		return nil, err
	}
	return article, err
}

// GetArticleByCategory genju
func GetArticleByCategory(categoryID int64, start, num int) (articleList []*model.ArticleInfo, err error) {
	sqlStr := "select id,category_id,title,summary,view_count,comment_count,user_name,create_time,updata_time from article where category_id = ? order by create_time desc limit ?,?"
	rows, err := DB.Query(sqlStr, categoryID, start, num)
	if err != nil {
		return
	}
	for rows.Next() {
		var myarticle model.ArticleInfo
		rows.Scan(&myarticle.ArticleID, &myarticle.ArticleCategory, &myarticle.ArticleTitle, &myarticle.ArticleSummary, &myarticle.ArticleViewCount, &myarticle.ArticleCommentCount, &myarticle.ArticleUserName, &myarticle.ArticleCreateTime, &myarticle.ArticleUpdataTime)
		articleList = append(articleList, &myarticle)
	}
	return
}
