package model

// `id` `category_name` `category_no` `create_time` `updata_time`

// Category 标签结构体
type Category struct {
	CategoryID   int64  `db:"id"`
	CategoryName string `db:"category_name"`
	CategoryNo   int    `db:"category_no"`
}
