package helpers

import (
	"encoding/json"
)

// StringToStruct 将字符串转换为结构体
func StringToStruct(result string, target interface{}) error {
	if err := json.Unmarshal([]byte(result), &target); err != nil {
		// log.Logf("error in StringToStruct %v", err)
		return err
	}

	return nil
}

// JSONStringify 将对象转换为json字符串数据
func JSONStringify(data interface{}) (string, error) {
	str, err := json.Marshal(data)
	if err != nil {
		// log.Logf("error JSONStringify %v", err)
		return "", err
	}

	return string(str), nil
}
