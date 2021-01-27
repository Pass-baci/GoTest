package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转整型
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)	返回有符号的整型
	//base代表进制，base 的值："0x" 表示 16 进制； "0" 表示 8 进制；否则就是 10 进制。
	//bitSize 表示的是整数取值范围，或者说整数的具体类型。取值 0、8、16、32 和 64 分别代表 int、int8、int16、int32 和 int64。
	// func ParseUint(s string, base int, bitSize int) (n uint64, err error)	返回无符号的整型
	// func Atoi(s string) (i int, err error)	返回int类型的整型
	parseInt, err := strconv.ParseInt("127", 10, 8)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parseInt)
	//运行结果
	//strconv.ParseInt: parsing "128": value out of range
	//127

	//整型转为字符串
	//func FormatUint(i uint64, base int) string    // 无符号整型转字符串
	//func FormatInt(i int64, base int) string    // 有符号整型转字符串
	//func Itoa(i int) string
	u := uint64(127)
	i := int(-127)
	formatUint := strconv.FormatUint(u, 10)
	itoa := strconv.Itoa(i)
	fmt.Println(formatUint,itoa)
}
