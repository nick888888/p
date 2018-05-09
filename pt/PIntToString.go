package pt

import (
	"strconv"
	"fmt"
	"reflect"
)

func IntToString(){
	var (
		s string
		x = 100
	)
	s = strconv.Itoa(x) //将一个整数转换成字符串的方法
	fmt.Println("转换前的类型",reflect.TypeOf(x))
	fmt.Println("我是转换后的字符串：",s)
	fmt.Println("转换后的类型",reflect.TypeOf(s))

}

//#以上代码输出结果如下：
//转换前的类型 int
//我是转换后的字符串： 100
//转换后的类型 string