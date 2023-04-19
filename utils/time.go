package utils

import "time"

/**
* @Author: Xenolies
* @Date: 2023/4/18 22:18
* @Version: 1.0
 */

/*UnixTimeConvert Unix时间转换为String
* UnixTime Unix时间 DateFormat 日期格式
 */
func UnixTimeConvert(UnixTime int64, DateFormat string) string {
	//返回time对象
	t := time.Unix(UnixTime, 0)

	//返回string
	dateStr := t.Format(DateFormat)

	return dateStr
}
