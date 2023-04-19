package controller

import (
	"bolg_go/database"
	"bolg_go/middleware"
	"bolg_go/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
* @Author: Xenolies
* @Date: 2023/4/16 12:10
* @Version: 1.0
 */

// Checkbox 复选框条件
type Checkbox struct {
	Conditions []string `form:"conditions[]"`
}

// NotFound 设置默认路由当访问一个错误网站时返回
func NotFound(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "404.html", 0)
}

//func AdminLogin(ctx *gin.Context) {
//	userName := ctx.PostForm("username")     // 用户名
//	userPassword := ctx.PostForm("password") // 用户密码
//	user := model.Users{}
//
//	database.DefaultDb.Table("users").Where("user_name =  ? ", userName).First(&user)
//	if user.Password == userPassword && len(userName) > 0 && len(userPassword) > 0 {
//		ctx.JSON(http.StatusOK, gin.H{
//			"status": 200,
//			"msg":    "Login 请求成功! ",
//		})
//		token, err := middleware.MakeToken(userName, userPassword) // 生成token
//		if err != nil {
//			errors.New("生成Token失败")
//		}
//		ctx.Header("token", token) //在header里添加token
//	} else {
//		ctx.Abort()
//	}
//
//}

// Login 登录
// TODO: 需要优化,可能和 GetOneUser 冲突
func Login(ctx *gin.Context) {
	userName := ctx.PostForm("username")     // 用户名
	userPassword := ctx.PostForm("password") // 用户密码
	user := model.Users{}

	sfID := middleware.GetSnowflakeIDInt64() // 生成的雪花int64 ID

	// select user_name, password from users where user_name = userName;
	database.DefaultDb.Table("users").Select("user_name , password , mail").Where("user_name =  ? ", userName).First(&user)
	if user.Password == userPassword {

		token, err := middleware.MakeToken(user) // 生成token
		if err != nil {
			errors.New("生成Token失败")
		}
		ctx.Header("token", token) //在header里添加token

		// 将用户token存入redis
		middleware.RedisDb.Set("user:"+strconv.FormatInt(sfID, 10), token, 300)

		ctx.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "Login 请求成功! ",
		})

	}
}

// Register 注册功能
// TODO: 需要优化,可能和 CreateOneUser 冲突
func Register(ctx *gin.Context) {
	userName := ctx.PostForm("username")     // 用户名
	userPassword := ctx.PostForm("password") // 用户密码
	userMail := ctx.PostForm("mail")         // 邮箱
	sfID := middleware.GetSnowflakeIDInt64()
	nickName := ctx.PostForm("nickname")

	// 如果昵称存在就用用户的
	if nickName == "" {
		nickName = "user_" + strconv.FormatInt(sfID, 10)
	}

	user := model.Users{
		ID:       sfID,
		NickName: nickName, // 默认用户名
		UserName: userName,
		Password: userPassword,
		Mail:     userMail,
		Level:    3, // 默认注册为普通用户
		Content:  "",
		Avatar:   "",
	}

	// 生成Token字段 使用 strconv.FormatInt 将int64 转为字符串
	token, err := middleware.MakeToken(user) // 生成token
	if err != nil {
		errors.New("生成Token失败")
	}
	ctx.Header("token", token) //在header里添加token

	database.DefaultDb.Create(&user)
	fmt.Println("创建用户: ", user.NickName)
	// 同时将用户信息存到Redis中
	middleware.RedisDb.Set("user:"+strconv.FormatInt(sfID, 10), token, 300)

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Register 请求成功! ",
	})

}

//// CreateOneUser 创建用户
//// TODO: 创建用户,可能和CreateOneUser冲突
//func CreateOneUser(ctx *gin.Context) {
//	//username :=  ctx.PostForm("username")
//	//database.DefaultDb.Table("users").Where("name = ?",username).First(&user.User{})
//	//ctx.JSON(http.StatusOK, gin.H{
//	//	"status": 200,
//	//	"msg":    "CreateOneUser 请求成功! ",
//	//})
//}

// UpdateUserByID 根据用户ID修改
func UpdateUserByID(ctx *gin.Context) {
	//TODO: 晚上接着写这段
	id := ctx.PostForm("id")

	user := model.Users{}

	// 查询 id 匹配字段你的所有信息
	database.DefaultDb.Where("id = ", id).First(&user)

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "UpdateOneUser 请求成功! ",
	})
}

// UpdateUserByConditions 根据用户ID修改
func UpdateUserByConditions(ctx *gin.Context) {
	//TODO: 逻辑没想好
	id := ctx.PostForm("id")
	user := model.Users{}

	// 查询 id 匹配字段你的所有信息
	database.DefaultDb.Where("id = ", id).First(&user)

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "UpdateOneUser 请求成功! ",
	})
}

func DeleteOneUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "DeleteOneUser 请求成功! ",
	})
}

func GetAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "GetAllUsers 请求成功! ",
	})
}

func FindUsers(ctx *gin.Context) {
	// 获取筛选条件
	var checkBox Checkbox
	ctx.ShouldBind(&checkBox)
	ctx.JSON(http.StatusOK, gin.H{
		"status":    200,
		"msg":       "FindUsers 请求成功! ",
		"condition": checkBox.Conditions,
	})
}

func GetOneUser(ctx *gin.Context) {
	userName := ctx.PostForm("username")     // 用户名
	userPassword := ctx.PostForm("password") // 用户密码
	user := model.Users{}

	sfID := middleware.GetSnowflakeIDInt64() // 生成的雪花int64 ID

	// select user_name, password from users where user_name = userName;
	database.DefaultDb.Table("users").Select("user_name , password , mail").Where("user_name =  ? ", userName).First(&user)
	if user.Password == userPassword {

		token, err := middleware.MakeToken(user) // 生成token
		if err != nil {
			errors.New("生成Token失败")
		}
		ctx.Header("token", token) //在header里添加token

		// 将用户token存入redis
		//TODO: 这里的逻辑需要优化,strconv.FormatInt() 可能一套写个封装作为专门转换
		middleware.RedisDb.Set("user:"+strconv.FormatInt(sfID, 10), token, 300)

	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "GetOneUser 请求成功! ",
	})
}

// AddOneUser 添加用户
func AddOneUser(ctx *gin.Context) {
	userName := ctx.PostForm("username")     // 用户名
	userPassword := ctx.PostForm("password") // 用户密码
	userMail := ctx.PostForm("mail")         // 邮箱
	sfID := middleware.GetSnowflakeIDInt64()
	nickName := ctx.PostForm("nickname")

	// 如果昵称存在就用用户的
	if nickName == "" {
		nickName = "user_" + strconv.FormatInt(sfID, 10)
	}

	user := model.Users{
		ID:       sfID,
		NickName: nickName, // 默认用户名
		UserName: userName,
		Password: userPassword,
		Mail:     userMail,
		Level:    3, // 默认注册为普通用户
		Content:  "",
		Avatar:   "",
	}

	// 生成Token字段 使用 strconv.FormatInt 将int64 转为字符串
	token, err := middleware.MakeToken(user) // 生成token
	if err != nil {
		errors.New("生成Token失败")
	}
	ctx.Header("token", token) //在header里添加token

	database.DefaultDb.Create(&user)
	fmt.Println("创建用户: ", user.NickName)
	// 同时将用户信息存到Redis中
	middleware.RedisDb.Set("user:"+strconv.FormatInt(sfID, 10), token, 300)

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "AddOneUser 请求成功! ",
	})

}
