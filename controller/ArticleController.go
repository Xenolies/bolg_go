package controller

import (
	"bolg_go/database"
	"bolg_go/model"
	"bolg_go/utils"
	"github.com/gin-gonic/gin"
	"time"
)

/**
* @Author: Xenolies
* @Date: 2023/4/18 10:12
* @Version: 1.0
 */

// PostArticles 发送文章上传到数据库
func PostArticles(ctx *gin.Context) {
	articleName := ctx.PostForm("article_name")  // 文章名
	introduction := ctx.PostForm("introduction") // 文章简介
	context := ctx.PostForm("context")           // 文章内容

	// 如果 简介为空 就截取文章内容的前100个字符
	// 如果有就设置简介
	if introduction == "" {

		introduction = utils.TruncatedString(introduction)
	}

	arts := model.Articles{
		ID:                  100,
		ArticleName:         articleName,
		ArticleIntroduction: introduction,
		Content:             context,
		Cover:               "",
		CreatedAt:           time.Time{},
		UpdatedAt:           time.Time{},
	}
	database.DefaultDb.Create(&arts)

}
