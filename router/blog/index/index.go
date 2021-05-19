package blogindex

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router(r *gin.Engine) {
	//// 通过内置的 html/template 实现
	//r.GET("/", func(c *gin.Context) {
	//	c.Status(http.StatusOK)
	//	tpl, err := template.ParseFiles("template/blog/index/index.html")
	//	if err != nil {
	//		log.Printf("Parse template failed: %v", err)
	//		return
	//	}
	//
	//	name := "ABC"
	//	err = tpl.Execute(c.Writer, name)
	//	if err != nil {
	//		log.Printf("Render template failed: %v", err)
	//		return
	//	}
	//
	//	c.HTML(http.StatusOK, "index.html", "2021")
	//})

	// 通过 gin 框架实现
	r.LoadHTMLGlob("template/blog/index/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "2021")
	})
}