package util

import (
	"github.com/daobin/golang-training-camp/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPageRowStart(c *gin.Context) int {
	rowStart := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		rowStart = (page - 1) * setting.AppSetting.PageSize
	}

	return rowStart
}
