package utils

import "os"

/**
* @Author: Xenolies
* @Date: 2023/4/18 22:17
* @Version: 1.0
 */

// PathExists 判断路径（包括文件与文件夹）是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
