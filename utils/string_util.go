package utils

import (
	"fmt"
	"strings"
)



type _stringUtil struct {

}
const(
	StringNone=iota //AbC -> AbC
	StringLower //AbC -> abc
	StringUpper //AbC -> ABC
	StringFirstLetterLowerNotAddLineOtherLowerAddLine //AbCD -> abc_d
	StringLowerAddLine //AbCD -> a_bc_d
	StringFirstLetterUpperNotAddLineOtherUpperAddLine //AbCD -> AbCD abc_d -> abCd
	StringUpperAddLine //AbCD -> AbCD a_bc_d -> AbCd
)
func (*_stringUtil)IsBank(str string)  bool{
	return str==""||strings.Trim(str," ")==""
}
func (*_stringUtil)Parse(str string,stringFlag int) string {
	switch stringFlag {
		case StringNone:return  str

		case StringLower:return strings.ToLower(str)

		case StringUpper:return strings.ToUpper(str)

		case StringFirstLetterLowerNotAddLineOtherLowerAddLine:
			{
				var b strings.Builder
				b.Grow(len(str))
				for i := 0; i < len(str); i++ {
					c := str[i]
					if '_'==c{
						continue
					}
					if 'A'<=c&&c<='Z'{
						c += 'a' - 'A'
						if i!=0{
							b.WriteByte('_')
						}
						b.WriteByte(c)

					}else{
						b.WriteByte(c)
					}
				}
				return  b.String()
			}

		case StringLowerAddLine:{
			var b strings.Builder
			b.Grow(len(str))
			for i := 0; i < len(str); i++ {
				c := str[i]
				if 'A'<=c&&c<='Z'{
					c += 'a' - 'A'
					b.WriteByte('_')
					b.WriteByte(c)

				}else{
					b.WriteByte(c)
				}
			}
			return  b.String()
		}

		case StringFirstLetterUpperNotAddLineOtherUpperAddLine:{
			var b strings.Builder
			b.Grow(len(str))
			for i := 0; i < len(str); i++ {
				c := str[i]
				if   'a'<=c&&c<='z'{
					if i==0|| i+1<len(str)&&str[i+1]=='_'{
						c -= 'a' - 'A'
						b.WriteByte(c)
					}else{
						b.WriteByte(c)
					}
				}else if c=='_'{
					continue
				}else{
					b.WriteByte(c)
				}
			}
			return  b.String()
		}

		case StringUpperAddLine:
			var b strings.Builder
			b.Grow(len(str))
			for i := 0; i < len(str); i++ {
				c := str[i]
				if   'a'<=c&&c<='z'{
					if i+1<len(str)&&str[i+1]=='_'{
						c -= 'a' - 'A'
						b.WriteByte(c)
					}else{
						b.WriteByte(c)
					}
				}else if c=='_'{
					continue
				}else{
					b.WriteByte(c)
				}
			}
			return  b.String()

	}

	return str
}
//https://www.cnblogs.com/Detector/p/9686443.html
// Capitalize 字符首字母大写
func (*_stringUtil)Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
//a_b_c false A_bc true
func (*_stringUtil)IsFirstLetterUpper(str string) bool {
	vv := []rune(str)   // 后文有介绍
	if vv[0] >= 65 && vv[0] <= 90{return true}
	return  false
}
//a_b_c true A_bc false
func (*_stringUtil)IsFirstLetterLower(str string) bool {
	vv := []rune(str)   // 后文有介绍
	if vv[0] >= 97 && vv[0] <= 122{return true}
	return  false
}
//a_b_c ABC
func (*_stringUtil)LowerToUpper(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else  if vv[i]=='_'{
				continue
			} else {
				upperStr += string(vv[i])
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
// ABC a_b_c
func (*_stringUtil)UpperToLower(str string) string {
	var lowerStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i]=='_'{
				continue
			} else if vv[i] >= 65 && vv[i] <= 90{
				vv[i] += 32 // string的码表相差32位
				lowerStr += string(vv[i])+"_"
			}else {
				lowerStr += string(vv[i])
			}
		} else {
			lowerStr += string(vv[i])
		}
	}
	return lowerStr
}
func (*_stringUtil)Table(str string) string {
	var lowerStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i]=='_'{
				continue
			} else if vv[i] >= 65 && vv[i] <= 90{
				vv[i] += 32 // string的码表相差32位
				lowerStr += string(vv[i])+"_"
			}else {
				lowerStr += string(vv[i])
			}
		} else {
			lowerStr += string(vv[i])
		}
	}
	return lowerStr
}
func(*_stringUtil) IndexOf(str string, ch byte) int{
	if str == "" {
		return -1
	}
	var chars = []byte(str)
	for i := range chars {
		if str[i] == ch {
			return int(i)
		}
	}
	return -1
}

func(*_stringUtil) Replace(str string) string {
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\b", "")
	return str
}
func(*_stringUtil) ReplaceModel(str string) string {
	str = strings.ReplaceAll(str, "Model", "")
	str = strings.ReplaceAll(str, "Bean", "")
	str = strings.ReplaceAll(str, "Entry", "")
	str = strings.ReplaceAll(str, "Entity", "")
	return str
}
func(*_stringUtil)  Substring(str string, index int) string {
	if str == "" {
		return ""
	}
	var chars = []byte(str)
	var ll = len(str)
	var l = ll
	if ll > index {
		l = index
	} else {
		l = ll
	}
	var temps = make([]byte, l)
	for i := 0; i < l; i++ {
		temps[i] = chars[i]
	}
	return string(temps)
}


