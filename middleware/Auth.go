package middleware

/**
* @Author: Xenolies
* @Date: 2023/4/17 9:39
* @Version: 1.0
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/**
中间件
*/

// AuthMiddleware 封包认证
// TODO:  路径判断需要优化
func AuthMiddleware(doCheck bool, ReleaseStructure ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if doCheck {
			switch strings.ToLower(ctx.Request.RequestURI) {
			case "/user/login":
				ctx.Next()

			case "/admin/login":
				ctx.Next()

			case "/user/register":
				ctx.Next()

			default:
				token := ctx.Request.Header.Get("token")
				if token == "" {
					ctx.JSON(http.StatusOK, gin.H{
						"status": -1,
						"msg":    "请求未携带token,无权限访问",
					})
					ctx.Abort()
					return
				} else {
					_, err := ParseToken(token)
					if err != nil {
						fmt.Println("Token解码失败或过期: ", err)
						ctx.JSON(http.StatusOK, gin.H{
							"status": -1,
							"msg":    "Token已过期,请重新登录",
						})
						// 截止后续请求
						ctx.Abort()
						return
					}
				}
			}

			//过滤是否验证token， login 和 register 结构直接放行，这里为了简单起见，直接判断路径中是否带login，携带login直接放行
			//if strings.Contains(ctx.Request.RequestURI, "login") {
			//	return
			//}
			//if strings.Contains(ctx.Request.RequestURI, "register") {
			//	return
			//}
			//放行结构
			//if strings.ToLower(ctx.Request.RequestURI) ==

			//token := ctx.Request.Header.Get("token")
			//if token == "" {
			//	ctx.JSON(http.StatusOK, gin.H{
			//		"status": -1,
			//		"msg":    "请求未携带token,无权限访问",
			//	})
			//	ctx.Abort()
			//	return
			//} else {
			//	_, err := ParseToken(token)
			//	if err != nil {
			//		fmt.Println("Token解码失败或过期: ", err)
			//		//ctx.JSON(http.StatusOK, gin.H{
			//		//	"status": -1,
			//		//	"msg":    "Token已过期,请重新登录",
			//		//})
			//		// 截止后续请求
			//		ctx.Redirect(http.StatusOK, "/login")
			//	}
			//}

		} else {
			// 不使用直接放行
			ctx.Next()
		}
	}
}
