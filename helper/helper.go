package helper

import (
	"regexp"
	"strconv"
)

func CheckIdCard(idCardStr string) bool {
	//18位身份证 ^(\d{17})([0-9]|X)$
	// 匹配规则
	// (^\d{15}$) 15位身份证
	// (^\d{18}$) 18位身份证
	// (^\d{17}(\d|X|x)$) 18位身份证 最后一位为X的用户
	regRuler := "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)

	// 返回 MatchString 是否匹配
	if !reg.MatchString(idCardStr) {
		return false
	}

	idCardBytes := []byte(idCardStr)

	// 通过前17位计算最后一位
	array := make([]int, 17)

	// 强制类型转换，将[]byte转换成[]int ,变化过程
	// []byte -> byte -> string -> int
	// 将通过range 将[]byte转换成单个byte,再用强制类型转换string()，将byte转换成string
	// 再通过strconv.Atoi()将string 转换成int 类型
	for index, value := range idCardBytes[0:17] {
		array[index], _ = strconv.Atoi(string(value))
	}

	var weight = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var res int
	for i := 0; i < 17; i++ {
		res += array[i] * weight[i]
	}

	lastByte := res % 11
	a18 := [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	if a18[lastByte] == idCardBytes[len(idCardBytes)-1] {
		return true
	}
	return false
}
