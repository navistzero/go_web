package db

import (
	"fmt"
	"obedu/blogger/model"
	"strings"
)

// InsertCategory ...
func InsertCategory(category *model.Category) (categoryID int64, err error) {
	sqlStr := "insert into category(category_name,category_no) value(?,?)"
	// tx, err := DB.Begin()
	// if nil != err {
	// 	fmt.Println("开启事务失败:", err)
	// 	return -1, err
	// }
	stmt, err := DB.Prepare(sqlStr)
	if nil != err {
		fmt.Println("Prepare Failed!")
		return -1, err
	}
	// 将参数传递到SQL语句中执行
	res, err := stmt.Exec(category.CategoryName, category.CategoryNo)
	if nil != err {
		fmt.Println("Exec Failed!")
		return -1, err
	}
	// 将事务提交
	// tx.Commit()
	categoryID, err = res.LastInsertId()
	return
}

// GetCategoryByID 单个
func GetCategoryByID(id int64) (category *model.Category, err error) {
	sqlStr := "select id,category_name,category_no from category where id = ?"
	category = &model.Category{}
	err = DB.QueryRow(sqlStr, id).Scan(&category.CategoryID, &category.CategoryName, &category.CategoryNo)
	if err != nil {
		return nil, err
	}
	return category, err
}

// GetCategoryList 根据id获取多个值，
func GetCategoryList(categorIDs []string) (categorys []*model.Category, err error) {
	sqlStr := "select id,category_name,category_no from category where id in ("
	sqlStr += strings.Join(categorIDs, ",") + ")"
	// for _, val := range categorIDs {
	// 	sqlStr += strconv.FormatInt(val, 10) + ","
	// }
	// nsqlStr := sqlStr[:len(sqlStr)-1] + ")"
	rows, err := DB.Query(sqlStr)
	if err != nil {
		return
	}
	for rows.Next() {
		var category model.Category
		rows.Scan(&category.CategoryID, &category.CategoryName, &category.CategoryNo)
		categorys = append(categorys, &category)
	}
	return
}

// GetAllCategory 获取所有
func GetAllCategory() (categorys []*model.Category, err error) {
	sqlStr := "select id,category_name,category_no from category"
	rows, err := DB.Query(sqlStr)
	if err != nil {
		return
	}
	for rows.Next() {
		var category model.Category
		rows.Scan(&category.CategoryID, &category.CategoryName, &category.CategoryNo)
		categorys = append(categorys, &category)
	}
	return
}
