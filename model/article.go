package model

import (
	"time"
)

// `id` `content` `category_id` `title` `view_count` `comment_count` `user_name` `status` `summary` `create_time` `updata_time`

// ArticleInfo 文章结构体
type ArticleInfo struct {
	ArticleID           int64     `db:"id"`
	ArticleCategory     int64     `db:"category_id"`
	ArticleTitle        string    `db:"title"`
	ArticleSummary      string    `db:"summary"`
	ArticleViewCount    int       `db:"view_count"`
	ArticleCommentCount int       `db:"comment_count"`
	ArticleUserName     string    `db:"user_name"`
	ArticleCreateTime   time.Time `db:"create_time"`
	ArticleUpdataTime   time.Time `db:"updata_time"`
}

//ArticleDetail 文章详情页
type ArticleDetail struct {
	ArticleInfo
	ArticleContent string `db:"content"`
	Category
}

// ArticleRecord 文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}
