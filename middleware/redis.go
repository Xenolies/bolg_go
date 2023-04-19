package middleware

/**
* @Author: Xenolies
* @Date: 2023/4/17 22:37
* @Version: 1.0
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

/*
redis 数据库中间件
*/

// RedisDb 全局redis变量
var RedisDb *redis.Client

// 初始化 建立redis链接
func init() {
	//ctx:= context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := RedisDb.Ping().Result() // 测试链接
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
	fmt.Println("Redis数据库链接成功")
}

func getUserTokenFromRedis(ctx *gin.Context) {
	// user:userid value :
	//tokenHandler := ctx.Request.Header.Get("token") // 获取在handler 中保存的token字段
	//userName := ctx.PostForm("username ")

}
