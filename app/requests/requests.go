package requests

import (
	"go_simpleweibo/pkg/utils"
	"regexp"
	"strconv"
	"strings"
)

type (
	// 验证器函数
	ValidatorFunc = func() (msg string)
	// 验证器数组 map
	ValidatorMap = map[string][]ValidatorFunc
	// 错误信息数组
	ValidatorMsgArr = map[string][]string
)

func RunValidators(m ValidatorMap, msgMap ValidatorMsgArr) (errors []string) {

	for k, validators := range m {
		customMsgArr := msgMap[k] // 自定义错误信息数组
		customMsgArrLen := len(customMsgArr)

		for i, fn := range validators {
			msg := fn()
			if msg != "" {
				if i < customMsgArrLen && customMsgArr[i] != "" {
					// 采用自定义的错误信息输出
					msg = customMsgArr[i]
				} else {
					// 采用默认的错误信息输出
					names := strings.Split(k, "|")
					data := make(map[string]string)

					for ti, tv := range names {
						data["$key"+strconv.Itoa(ti+1)+"$"] = tv
					}

					msg = utils.ParseEasyTemplate(msg, data)
				}

				errors = append(errors, msg)
				break // 进行下一个字段的验证
			}
		}
	}

	return errors
}

// RequiredValidator : value 必须存在
func RequiredValidator(value string) ValidatorFunc {
	return func() (msg string) {
		if value == "" {
			return "$key1$ 必须存在"
		}

		return ""
	}
}

// MixLengthValidator -
func MixLengthValidator(value string, minStrLen int) ValidatorFunc {
	return func() (msg string) {
		l := len(value)

		if l < minStrLen {
			return "$key1$ 必须大于 " + strconv.Itoa(minStrLen)
		}

		return ""
	}
}

// MaxLengthValidator -
func MaxLengthValidator(value string, maxStrLen int) ValidatorFunc {
	return func() (msg string) {
		l := len(value)

		if l > maxStrLen {
			return "$key1$ 必须小于 " + strconv.Itoa(maxStrLen)
		}

		return ""
	}
}

// EqualValidator -
func EqualValidator(v1 string, v2 string) ValidatorFunc {
	return func() (msg string) {
		if v1 != v2 {
			return "$key1$ 必须等于 $key2$"
		}

		return ""
	}
}

// EmailValidator 验证邮箱格式
func EmailValidator(value string) ValidatorFunc {
	return func() (msg string) {
		pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
		reg := regexp.MustCompile(pattern)
		status := reg.MatchString(value)

		if !status {
			return "$key1$ 邮箱格式错误"
		}

		return ""
	}
}
