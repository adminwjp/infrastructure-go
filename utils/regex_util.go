package utils

import (
	"regexp"
)

type _regexUtil struct {

}

func (reg _regexUtil) IsEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	return reg.IsMatch(email,pattern)
}
func (_regexUtil) IsMatch(str string,pattern string) bool {
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(str)
}
func (reg  _regexUtil) IsPhone(phone string) bool {
	pattern := `[13|15|17|18]\d{9}`
	return reg.IsMatch(phone,pattern)
	//match, _ := regexp.Match("[13|15|17|18]\\d{9}", []byte(user.Phone))
	//return  match
}