package main

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		//在golang中规定了switch type的case T1，类型列表只有一个
		//如果是case T1, T2，类型列表中有多个，那v的类型还是多对应接口的类型，也就是m的类型
		//msg的类型为interface{}
		msg.Name  //Unresolved reference 'Name'
	}
}
