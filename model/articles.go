package model

import "time"

/**
* @Author: Xenolies
* @Date: 2023/4/17 10:49
* @Version: 1.0
 */

// Articles 文章对象
type Articles struct {
	ID                  int64     `grom:"primarykey"`           // 文章ID
	ArticleName         string    `gorm:"article_name"`         // 文章名称
	ArticleIntroduction string    `gorm:"article_introduction"` // 简介
	Content             string    `gorm:"content"`              //内容
	Cover               string    `gorm:"cover"`                // 封面
	CreatedAt           time.Time `gorm:"create_at"`            // 创建时间
	UpdatedAt           time.Time `gorm:"update_at"`            // 更新时间
}
