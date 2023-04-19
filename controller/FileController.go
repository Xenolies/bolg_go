package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* @Author: Xenolies
* @Date: 2023/4/18 10:14
* @Version: 1.0
 */

// UploadOneFile 上传单个文件
func UploadOneFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败"})
		return
	}
	// 保存文件到本地
	fmt.Println("文件名为: ", file.Filename) //输出文件名
	err = ctx.SaveUploadedFile(file, file.Filename)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("保存上传文件失败 %s", err.Error()))
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf("上传文件 %s 成功", file.Filename))
}

// UploadFiles 多文件上传
func UploadFiles(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files, _ := form.File["upload[]"] //获取文件数组

	for _, file := range files { //遍历保存
		//获取文件名
		fileName := file.Filename
		fmt.Println("文件名为: ", fileName) //输出文件名

		//将上传的文件保存带本地
		err := ctx.SaveUploadedFile(file, fileName)
		if err != nil {
			ctx.String(http.StatusBadRequest, "文件保存失败 Err: %s", err.Error())
			return
		}

		ctx.String(http.StatusOK, "%s uploaded!\n", file.Filename)
	}
}
