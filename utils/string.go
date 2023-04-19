package utils

/**
* @Author: Xenolies
* @Date: 2023/4/18 22:17
* @Version: 1.0
 */

// TruncatedString 截取前100个字符
func TruncatedString(introString string) string {
	rs := []rune(introString)
	return string(rs[0:99])
}
