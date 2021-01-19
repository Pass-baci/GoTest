package split

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a","b","c"}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("excepted:%#v, got:%#v", want, got)
	}
}
//$ go test
//输出：
//PASS
//ok      github.com/pass_baci/GoTest/20210118/单元测试/split     0.018s

func TestSplit2(t *testing.T) {
	got := Split("abcd","bc")
	want := []string{"a","d"}
	if !reflect.DeepEqual(want,got) {
		t.Errorf("excepted:%#v, got:%#v", want, got)
	}
}

//--- FAIL: TestSplit2 (0.00s)
//    split_test.go:24: excepted:[a d], got:[a cd]
//FAIL
//exit status 1
//FAIL    github.com/pass_baci/GoTest/20210118/单元测试/split     0.017s

//go test -v(查看测试函数名称和运行时间)
//=== RUN   TestSplit
//--- PASS: TestSplit (0.00s)
//=== RUN   TestSplit2
//    split_test.go:24: excepted:[a d], got:[a cd]
//--- FAIL: TestSplit2 (0.00s)
//FAIL
//exit status 1
//FAIL    github.com/pass_baci/GoTest/20210118/单元测试/split     0.018s

//定义一个测试组-示例
func TestSplit3(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		str string
		sep string
		want []string
	}
	// 定义一个存储测试用例的切片
	tests := []test{
		{str: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{str: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{str: "abcd", sep: "bc", want: []string{"a", "d"}},
		{str: "沙河有沙又有河", sep: "沙", want: []string{"","河有", "又有河"}},
	}
	//使用for循环进行执行测试
	for _, tc := range tests {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got,tc.want) {
			t.Errorf("excepted:%#v, got:%#v", tc.want, got)
		}
	}
}

//=== RUN   TestSplit
//--- PASS: TestSplit (0.00s)
//=== RUN   TestSplit2
//--- PASS: TestSplit2 (0.00s)
//=== RUN   TestSplit3
//    split_test.go:63: excepted:[]string{"河有", "又有河"}, got:[]string{"", "河有", "又有河"}
//--- FAIL: TestSplit3 (0.00s)
//FAIL
//exit status 1
//FAIL    github.com/pass_baci/GoTest/20210118/单元测试/split     0.019s

//子测试-示例
func TestSplit4(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"","河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("-----------以下为基准函数--------------")
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}
//goos: windows
//goarch: amd64
//pkg: github.com/pass_baci/GoTest/20210118/单元测试/split
//BenchmarkSplit-4         4761632               263 ns/op
//说明：4表示GOMAXPROCS的值，即为CPU的核心数；4761632表示测试次数；263ns/op表示每次调用的耗时
//PASS
//ok      github.com/pass_baci/GoTest/20210118/单元测试/split     1.526s

//go test -bench=Split -benchmem
//goos: windows
//goarch: amd64
//pkg: github.com/pass_baci/GoTest/20210118/单元测试/split
//BenchmarkSplit-4         4724139               252 ns/op             112 B/op          3 allocs/op
//说明：112 B/op表示每次操作内存分配了112字节，3 allocs/op则表示每次操作进行了3次内存分配。
//PASS
//ok      github.com/pass_baci/GoTest/20210118/单元测试/split     1.477s