package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//字符串比较
	//func Compare(a, b string) int   字符串相等，返回0；如果 a 小于 b ，返回 -1 ，反之返回 1
	str1 := "abcsad"
	str2 := "hello"
	fmt.Println(strings.Compare(str1,str2))
	fmt.Println(str1 == str2) //比较直观
	//func EqualFold(s, t string) bool 计算s 与 t 忽略字母大小写后是否相等。
	fmt.Println(strings.EqualFold(str1,str2))
	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println("-----------以上为字符串比较----------")

	//是否存在某个字符或子串
	//func Contains(s, substr string) bool
	fmt.Println(strings.Contains("abc","a"))
	//func ContainsRune(s string, r rune) bool
	fmt.Println(strings.ContainsRune("abc",'a'))
	//func ContainsAny(s, chars string) bool  (任意一个字符（Unicode Code Point）如果在第一个参数 s 中存在，则返回 true)
	fmt.Println(strings.ContainsAny("abc","cd"))
	fmt.Println("-----------以上为是否存在某个字符或子串----------")

	//子串出现次数
	//func Count(s, sep string) int
	fmt.Println(strings.Count("aaaabcaa", ""))
	fmt.Println("-----------以上为是子串出现次数----------")

	//字符串分割
	//func Split(s, sep string) []string (会将sep去掉)
	//func SplitAfter(s, sep string) (保留sep)
	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))
	fmt.Println("-----------以上为字符串分割----------")

	//字符串是否有某个前缀或后缀
	//func HasPrefix(s, prefix string) bool (s 中是否以 prefix 开始)
	//func HasSuffix(s, suffix string) bool (s 中是否以 suffix 结尾)
	fmt.Println(strings.HasPrefix("Gopher", "go"))
	fmt.Println(strings.HasSuffix("migo", "go"))
	fmt.Println("-----------以上为字符串是否有某个前缀或后缀----------")

	//字符或子串在字符串中出现的位置
	//func Index(s, sep string) int (在 s 中查找 sep 的第一次出现，返回第一次出现的索引)
	//func LastIndex(s, sep string) int (在 s 中查找 sep 的最后一次出现，返回最后一次出现的索引)
	fmt.Println(strings.Index("Gopher", "go"))
	fmt.Println(strings.LastIndex("migo", "go"))
	fmt.Println("-----------以上为字符或子串在字符串中出现的位置----------")

	//字符串 JOIN 操作
	//func Join(a []string, sep string) string (将字符串切片的元素使用sep连接起来)
	str := []string{"hello", "world", "!"}
	fmt.Println(strings.Join(str,"!"))
	fmt.Println("-----------以上为字符串 JOIN 操作----------")

	//字符串重复几次
	//func Repeat(s string, count int) string
	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println("-----------以上为字符串重复几次----------")

	//字符替换
	//func Map(mapping func(rune) rune, s string) string
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
	fmt.Println("-----------以上为字符替换----------")

	//字符串子串替换
	//func Replace(s, old, new string, n int) string (用 new 替换 s 中的 old，一共替换 n 个)
	//func ReplaceAll(s, old, new string) string (全部替换)
	fmt.Println(strings.Replace("abbbbbbcdef", "b", "",5))
	fmt.Println(strings.ReplaceAll("abbbbbbcdef", "b",""))
	fmt.Println("-----------以上为字符串子串替换----------")

	//大小写转换
	//func ToLower(s string) string (转换成小写)
	//func ToLowerSpecial(c unicode.SpecialCase, s string) string (特殊字符转换成小写)
	//func ToUpper(s string) string (转换成大写)
	//func ToUpperSpecial(c unicode.SpecialCase, s string) string (特殊字符转换成大写)
	fmt.Println(strings.ToLower("HellO"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
	fmt.Println("-----------以上为大小写转换----------")

	//标题处理
	//func Title(s string) string (将首字母大写)
	//func ToTitle(s string) string (全部大写)
	//func ToTitleSpecial(c unicode.SpecialCase, s string) string ( 将 s 的每个字母大写，并且会将一些特殊字母转换为其对应的特殊大写字母)
	fmt.Println(strings.Title("hello wOrLd"))
	fmt.Println(strings.ToTitle("hello wOrLd"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "āáǎà ōóǒò êēéěè"))
	fmt.Println("-----------以上为标题处理----------")

	//修剪
	//func Trim(s string, cutset string) string
	fmt.Println(strings.Trim("abca","a"))
	fmt.Println("-----------以上为标题修剪----------")

	//Replacer 类型
	//func NewReplacer(oldnew ...string) *Replacer (进行多个替换)
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))
	fmt.Println("-----------以上为Replacer类型----------")

}
