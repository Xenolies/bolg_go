package model

import "time"

/**
* @Author: Xenolies
* @Date: 2023/4/17 10:49
* @Version: 1.0
 */

// Users 用户对象
type Users struct {
	ID        int64     `gorm:"primarykey"`
	NickName  string    `gorm:"nick_name"` //用户昵称
	UserName  string    `gorm:"user_name"` // 用户登录名
	Password  string    `gorm:"password"`  //登录密码
	Mail      string    `gorm:"mail"`      // 登录邮箱
	Level     int       `gorm:"level"`     // 用户权限
	Content   string    `gorm:"content"`   // 用户简介
	Avatar    string    `gorm:"avatar"`    // 头像
	CreatedAt time.Time `gorm:"create_at"` // 创建时间
	UpdatedAt time.Time `gorm:"update_at"` // 修改时间
}
