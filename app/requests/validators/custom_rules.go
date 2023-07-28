// Package validators 存放自定义规则和验证器
package validators

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gentleman/pkg/database"
	"net/http"
	"strings"

	"github.com/thedevsaddam/govalidator"
)

// 此方法会在初始化时执行，注册自定义表单验证规则
func init() {

	// 自定义规则 not_exists，验证请求数据必须不存在于数据库中。
	// 常用于保证数据库某个字段的值唯一，如用户名、邮箱、手机号、或者分类的名称。
	// not_exists 参数可以有两种，一种是 2 个参数，一种是 3 个参数：
	// not_exists:users,email 检查数据库表里是否存在同一条信息
	// not_exists:users,email,32 排除用户掉 id 为 32 的用户
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第一个参数，表名称，如 users
		tableName := rng[0]
		// 第二个参数，字段名称，如 email 或者 phone
		dbFiled := rng[1]

		// 第三个参数，排除 ID
		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}

		// 用户请求过来的数据
		requestValue := value.(string)

		// 拼接 SQL
		query := database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue)

		// 如果传参第三个参数，加上 SQL Where 过滤
		if len(exceptID) > 0 {
			query.Where("id != ?", exceptID)
		}

		// 查询数据库
		var count int64
		query.Count(&count)

		// 验证不通过，数据库能找到对应的数据
		if count != 0 {
			// 如果有自定义错误消息的话
			if message != "" {
				return errors.New(message)
			}
			// 默认的错误消息
			return fmt.Errorf("%v 已被占用", requestValue)
		}
		// 验证通过
		return nil
	})

	govalidator.AddCustomRule("check-lc", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)

		// 定义请求的URL和JSON数据
		url := "https://leetcode.cn/graphql/"

		payload := map[string]interface{}{
			"query": `query reputationUserReputations($userSlugs: [String!]!) {
			reputationUserReputations(userSlugs: $userSlugs) {
				level
				reputation
			}
		}`,
			"variables": map[string][]string{
				"userSlugs": {val},
			},
			"operationName": "reputationUserReputations",
		}

		// 转换JSON数据为字节流
		data, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling JSON data:", err)
			return nil
		}

		// 创建HTTP请求
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return nil
		}

		// 添加请求头信息
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Origin", "https://leetcode.cn")
		req.Header.Set("Referer", "https://leetcode.cn/u/gg_boy/")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.183")
		req.Header.Set("Authorization", "")
		req.Header.Set("Baggage", "sentry-environment=production,sentry-release=YekZwSfOpChArsaqNJG7J,sentry-transaction=%2Fu%2F%5Busername%5D,sentry-public_key=767ac77cf33a41e7832c778204c98c38,sentry-trace_id=b26d38716cef4875ac95ce4c881edb98,sentry-sample_rate=0.03")
		req.Header.Set("Random-UUID", "0aba5e34-df9b-9db9-fe45-496e73b25102")
		req.Header.Set("Sec-Ch-Ua", `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`)
		req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req.Header.Set("Sec-Ch-Ua-Platform", `"Windows"`)
		req.Header.Set("Sentry-Trace", "b26d38716cef4875ac95ce4c881edb98-9fbe3c36ae191f94-0")
		req.Header.Set("X-Csrftoken", "wWgiWBfqTrS5kBviDW8LmmHctGSXIBY6rZ0HpUrt51PnmLqnEtOeIouVJKFfDrSR")

		// 创建HTTP客户端并发送请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("Error sending request:%v", err)
		}
		defer resp.Body.Close()

		var result ResponseData

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return fmt.Errorf("Error decoding response:%v", err)
		}

		count := result.Data.ReputationUserReputations[0].Reputation

		if count == 0 {
			return fmt.Errorf("力扣用户名出错，请注意确认是用户名而不是昵称")
		}
		return nil

	})
}

// 定义结构体表示 JSON 数据的最外层
type ResponseData struct {
	Data Data `json:"data"`
}

// 定义结构体表示 "data" 字段
type Data struct {
	ReputationUserReputations []ReputationUser `json:"reputationUserReputations"`
}

// 定义结构体表示 "reputationUserReputations" 字段中的元素
type ReputationUser struct {
	Level      int `json:"level"`
	Reputation int `json:"reputation"`
}
