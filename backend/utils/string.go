package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mozillazg/go-pinyin"
	"io/ioutil"
	"strconv"
	"time"
)

func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

func StringToInt(e string) (int, error) {
	return strconv.Atoi(e)
}

func StringToUint(e string) (uint, error) {
	num, err := strconv.ParseUint(e, 10, 32)
	return uint(num), err
}

func StringToInt8(e string) (int8, error) {
	num, err := strconv.ParseInt(e, 10, 8)
	return int8(num), err
}

func GetCurrntTimeStr() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func GetCurrntTime() time.Time {
	return time.Now()
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

func GetBodyString(c *gin.Context) (string, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return string(body), nil
	} else {
		return "", err
	}
}

func JsonStrToMap(e string) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	} else {
		return nil, err
	}
}

func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

func ZhongwenToPinyin(s string) (result string) {
	// 创建一个Pinyin对象
	p := pinyin.NewArgs()
	p.Fallback = func(r rune, p pinyin.Args) []string {
		return []string{string(r)}
	}

	// 将中文转换为拼音
	pinyinResult := pinyin.Pinyin(s, p)

	// 将拼音结果连接起来
	for _, py := range pinyinResult {
		result += py[0]
	}
	return
}

// 生成随机密码
func GenerateRandomPassword(length int) (string, error) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.!@#$%&*"
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		indexBytes := make([]byte, 1)
		_, err := rand.Read(indexBytes)
		if err != nil {
			return "", err
		}

		index := int(indexBytes[0]) % len(charset)
		password[i] = charset[index]
	}

	return string(password), nil
}
