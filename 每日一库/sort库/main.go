package main

import (
	"fmt"
	"sort"
)

//sort包实现了四种基本排序算法：插入排序、归并排序、堆排序和快速排序，支持排序的数据类型有[]int,[]float64,[]string

//官网示例

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int 			{ return len(a) }  //集合内元素的总和

func (a ByAge) Swap(i, j int) 		{ a[i], a[j] = a[j], a[i] }  //交换下标 i j 对应的元素

func (a ByAge) Less(i, j int) bool 	{ return a[i].Age < a[j].Age }  //决定升序降序的条件

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	// 有两种方法可以对切片进行排序。首先，可以像ByAge一样为切片类型定义一组方法，然后调用sort.Sort。
	sort.Sort(ByAge(people))
	fmt.Println(people)
	//另一种方法是将sort.Slice与自定义的Less函数一起使用，该函数可以作为闭包提供。在这种情况下，不需要任何方法。 （如果存在，则将其忽略。）在这里，我们以相反的顺序重新排序：将闭包与ByAge.Less进行比较。
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)
	//sort包提供了 Reverse() 方法，可以允许将数据按 Less() 定义的排序方式逆序排序，而不必修改 Less() 代码
	sort.Sort(sort.Reverse(ByAge(people)))
	//sort包内除了排序，还具有查找功能（二分查找）
	//Search() 函数一个常用的使用方式是搜索元素是否在已经升序排好的切片中
	x := 11
	s := []int{3, 6, 8, 11, 45} // 注意已经升序排序
	pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
	if pos < len(s) && s[pos] == x {
		fmt.Println(x, " 在 s 中的位置为：", pos)
	} else {
		fmt.Println("s 不包含元素 ", x)
	}
}
//运行结果：
//[Bob: 31 John: 42 Michael: 17 Jenny: 26]
//[Michael: 17 Jenny: 26 Bob: 31 John: 42]
//[John: 42 Bob: 31 Jenny: 26 Michael: 17]

//基本数据类型的排序方法
//[]int排序常使用sort.Ints()，默认升序排序
//如果需要使用降序排序：sort.Sort(sort.Reverse(sort.IntSlice(s)))
