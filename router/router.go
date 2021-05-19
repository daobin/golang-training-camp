package router

import (
	"github.com/daobin/golang-training-camp/pkg/setting"
	blogindex "github.com/daobin/golang-training-camp/router/blog/index"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func staticFileRouter(r *gin.Engine) {
	r.StaticFS("/css", http.Dir("public/css"))
	r.StaticFS("/image", http.Dir("public/image"))
	r.StaticFS("/js", http.Dir("public/js"))
	r.StaticFS("/jquery", http.Dir("public/jquery"))
}

func InitRouter() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 静态资源路由
	staticFileRouter(r)

	// API 路由
	api := r.Group("/api")
	{
		api.GET("/tags")
		api.POST("/tags")
		api.PUT("/tags/:id")
		api.DELETE("/tags/:id")
	}

	// Blog 前台路由
	blogindex.Router(r)

	// Blog 后台路由
	//adm := r.Group("/adm")
	//{
	//	adm.GET("/")
	//}

	// 404 处理
	r.NoRoute(func(c *gin.Context) {
		msg := "噢，不！我居然迷路了！！"

		// 如果是 Ajax 或 api 请求，则返回 JSON
		xhr := strings.ToLower(c.Request.Header.Get("X-Requested-With"))
		reqUri := strings.ToLower(c.Request.RequestURI)
		if xhr == "xmlhttprequest" || strings.Contains(reqUri, "/api/") {
			c.JSON(http.StatusNotFound, gin.H{"msg": msg})
		} else {
			c.String(http.StatusNotFound, msg)
		}
	})

	return r
}