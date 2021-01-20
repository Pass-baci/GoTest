package main

import (
	"bytes"
	"fmt"
)

func main() {
	//字符串与字节间的转换
	str := "abc"
	b := []byte(str)
	str = string(b)
	fmt.Println(b, str)
	fmt.Println("-----------以上为字符串与字节间的转换------------")

	//func Contains(b, subslice []byte) bool (subslice切片是否在b切片上)
	isOk := bytes.Contains(b, []byte{97})
	fmt.Println(isOk)
	fmt.Println("-----------以上为是否存在某个子 slice------------")

	//[]byte 出现次数
	//func Count(s, sep []byte) int (sep在s中出现的次数)

	//Runes 类型转换
	//func Runes(s []byte) []rune
	runes := bytes.Runes(b)
	for _, v := range runes {
		fmt.Println(string(v))
	}
	fmt.Println("-----------以上为Runes 类型转换------------")

	//熟悉bytes.Reader
	x := []byte("你好，世界世界世界!!!") //定义一个字节切片
	reader := bytes.NewReader(x)	//实例化一个Reader对象
	d1 := make([]byte, len(x))		//初始化一个字节切片用于存储Reader对象将要读取的数据
	n, _ := reader.Read(d1)			//使用Read方法读取数据，返回字节数与错误
	fmt.Println(n, string(d1))

	reader1 := bytes.Reader{}
	reader1.Reset(x)
	d2 := make([]byte, len(x))
	n1, _ := reader1.Read(d2)
	fmt.Println(n1, string(d2))
	fmt.Println("-----------以上为bytes.Reader------------")

	//Reader的相关方法
	// 读取数据至 b
	//func (r *Reader) Read(b []byte) (n int, err error)
	//// 读取一个字节
	//func (r *Reader) ReadByte() (byte, error)
	readByte, _ := reader1.ReadByte()
	fmt.Println(readByte)
	//// 读取一个字符
	//func (r *Reader) ReadRune() (ch rune, size int, err error)
	//// 读取数据至 w
	//func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
	//// 进度下标指向前一个字节，如果 r.i <= 0 返回错误。
	//func (r *Reader) UnreadByte()
	//// 进度下标指向前一个字符，如果 r.i <= 0 返回错误，且只能在每次 ReadRune 方法后使用一次，否则返回错误。
	//func (r *Reader) UnreadRune()
	//// 读取 r.s[off:] 的数据至b，该方法忽略进度下标 i，不使用也不修改。
	//func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
	//// 根据 whence 的值，修改并返回进度下标 i ，当 whence == 0 ，进度下标修改为 off，当 whence == 1 ，进度下标修改为 i+off，当 whence == 2 ，进度下标修改为 len[s]+off.
	//// off 可以为负数，whence 的只能为 0，1，2，当 whence 为其他值或计算后的进度下标越界，则返回错误。
	//func (r *Reader) Seek(offset int64, whence int) (int64, error)

	//Buffer类型
	//Buffer类型的初始化方法
	bufferby := bytes.NewBuffer([]byte("Hello World"))
	bufferString := bytes.NewBufferString("Hello World")
	buffer := bytes.Buffer{}
	fmt.Println(bufferby, bufferString,buffer)

	//与Reader类型方法不同的三个方法
	// 读取到字节 delim 后，以字节数组的形式返回该字节及前面读取到的字节。如果遍历 b.buf 也找不到匹配的字节，则返回错误(一般是 EOF)
	//func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
	line, err := bufferString.ReadBytes('l')
	if err != nil {
		fmt.Println("delim:l,err:",err)
	} else {
		fmt.Println(string(line))
	}
	// 读取到字节 delim 后，以字符串的形式返回该字节及前面读取到的字节。如果遍历 b.buf 也找不到匹配的字节，则返回错误(一般是 EOF)
	//func (b *Buffer) ReadString(delim byte) (line string, err error)
	line1, err := bufferby.ReadString('l')
	if err != nil {
		fmt.Println("delim:l,err:",err)
	} else {
		fmt.Println(line1)
	}
	// 截断 b.buf , 舍弃 b.off+n 之后的数据。n == 0 时，调用 Reset 方法重置该对象，当 n 越界时（n < 0 || n > b.Len() ）方法会触发 panic.
	//func (b *Buffer) Truncate(n int)
	newBuffer := bytes.NewBuffer([]byte("Hello World"))
	newBuffer.Truncate(10)
	//newBuffer.WriteString("Good")
	fmt.Println(newBuffer.Len())
	readString, err := newBuffer.ReadString('W')
	if err != nil {
		fmt.Println("delim:H err:", err)
	} else {
		fmt.Println(readString)
	}
}
