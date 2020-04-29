package db

import (
	"fmt"
	"testing"
	"obedu/blogger/model"
)

func init() {
	err := Initdb()
	if err != nil {
		panic(err)
	}
}
func TestInsertCategory(t *testing.T) {
	cattest := model.Category{
		CategoryName: "python",
		CategoryNo:   1,
	}
	fmt.Println(cattest)
	cattestID, err := InsertCategory(&cattest)
	if err != nil {
		panic(err)
	}
	fmt.Println(cattestID)
}

// func TestGetCategoryByID(t *testing.T) {
// 	cat, err := GetCategoryByID(1)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(cat)
// }
// func TestGetCategoryList(t *testing.T) {
// 	catlist := []int64{1, 2, 3, 4}
// 	cat, err := GetCategoryList(catlist)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(cat)
// }
