package controller

import (
	"fmt"
	"net/http"
	server "obedu/blogger/service"

	"github.com/gin-gonic/gin"
)

// IndexPage shouye
func IndexPage(c *gin.Context) {
	articleList, err := server.GetArticleRecordList(0, 10)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	fmt.Println(articleList)
	c.HTML(200, "views/500.html", nil)

}
