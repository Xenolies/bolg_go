package middleware

import "github.com/bwmarrin/snowflake"

/**
* @Author: Xenolies
* @Date: 2023/4/18 11:10
* @Version: 1.0
 */

var node *snowflake.Node

func init() {
	node, _ = snowflake.NewNode(1)
}

func GetSnowflakeIDBase64() string {
	return node.Generate().Base64()
}

// GetSnowflakeIDInt64 生成唯一的雪花算法ID
// TODO: 总要 strconv.FormatInt() 可能要写个封装来将 int64 转为 String
func GetSnowflakeIDInt64() int64 {
	return node.Generate().Int64()
}
