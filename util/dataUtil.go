package util

import (
	"encoding/json"
	"time"
	"math/rand"
)

/**
json转map
 */
func Json2map(jsonStr string) (s map[string]string, err error) {
	var result map[string]string
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, err
	}
	return result, nil
}

/**
 * 字符串首字母转化为大写 ios_bbbbbbbb -> iosBbbbbbbbb
 */
func StrFirstToUpper(str string) string {
	var upperStr = ""
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			vv[i] -= 32
			upperStr += string(vv[i]) // + string(vv[i+1])
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

/**
随机生成字符串
 */
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
