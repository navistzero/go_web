package main

import (
	"fmt"

	"obedu/blogger/controller"
	"obedu/blogger/dao/db"

	"github.com/gin-gonic/gin"
	// server "obedu/blogger/service"
)

func main() {
	router := gin.Default()
	err := db.Initdb()
	if err != nil {
		fmt.Println(err)
	}
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("./views/*")
	router.GET("/", controller.IndexPage)
	router.Run(":8080")
	// cattest := model.Category{
	// 	CategoryName: "mysql",
	// 	CategoryNo:   3,
	// }
	// fmt.Println(cattest)
	// cattestID, err := db.InsertCategory(&cattest)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(cattestID)

	// cat, err := db.GetCategoryByID(2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(cat)

	// catlist := []string{"1", "2", "3"}
	// cat, err := db.GetCategoryList(catlist)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, ad := range cat {
	// 	fmt.Println(ad)
	// }

	// arttest1 := model.ArticleInfo{
	// 	ArticleCategory:     1,
	// 	ArticleTitle:        "title3",
	// 	ArticleSummary:      "summary3",
	// 	ArticleViewCount:    0,
	// 	ArticleCommentCount: 0,
	// 	ArticleUserName:     "lijun",
	// 	ArticleCreateTime:   time.Now(),
	// 	ArticleUpdataTime:   time.Now(),
	// }
	// arttest := model.ArticleDetail{
	// 	ArticleContent: "this is content three",
	// 	ArticleInfo:    arttest1,
	// }
	// fmt.Println(arttest)
	// cattestID, err := db.InsertArticle(&arttest)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(cattestID)

	// artlist, err := db.GetArticleList(0, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, ad1 := range artlist {
	// 	fmt.Println(ad1)
	// }

	// art, err := db.GetArticleByID(2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(art)

	// artlist, err := db.GetArticleByCategory(1, 0, 10)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, ad1 := range artlist {
	// 	fmt.Println(ad1)
	// }

	// artlist,err := server.GetArticleRecordList(0, 2)
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// for _,v := range artlist{
	// 	fmt.Println(v)
	// }

	// artlist,err := server.GetArticleRecordByCategory(1,0, 2)
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// for _,v := range artlist{
	// 	fmt.Println(v)
	// }
}
