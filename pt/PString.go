package pt

import (
	"fmt"
	"reflect"
	"strings"
)

//格式输出，定义多行注释
func StringA() {
	str1 := "yinzhengjie\thave a apple!" //"\t"表示制表符，会根据操作系统自动空出几个空格。
	doc := `  //定义多行注释
111
        222
        3333
        `
	x := "hello\\world" //转移用"／"
	y := "您好，尹正杰！\a"    //在终端执行电脑会响.在IDE上执行可能不会有声音哟！
	z := "####1\b2222"  //"\b"表示退格键，会把"\b"前面的一个字符给删除掉，即打印的时候会删除数字"1"
	fmt.Println(str1)
	fmt.Println(doc)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	//yinzhengjie	have a apple!
	////定义多行注释
	//111
	//   222
	//   3333
	//
	//hello\world
	//您好，尹正杰！
	//####2222

}

// 字符串修改
func StringB() {
	str1 := "yinzhengjie" + "world" //拼接
	fmt.Println(str1)

	var str2 byte     //1个字节（byte）=8个比特位（bits），因此byte = uint8
	str2 = str1[0]    //取字符
	fmt.Println(str2) //这里是字母y，但是输出是121
	//
	fmt.Printf("数字：%d 字符：%c\n", str2, str2) //表示字符"y"在ASCII表中对应的数字。printf不会换行但是println会自动换行。

	str3 := str1[0:11] //切片,顾左不顾右，取0-10共11个字符，不会去知道第12个字符，即不会取到下标为11第字符"w"
	fmt.Println(str3)  //yinzhengjie

	str4 := str1[11:len(str1)] //取索引11到最后
	fmt.Println(str4)          //world

	var b uint8 //我们这里也可以把unit8写成byte，其实是等效的。
	for b = 0; b < 177; b++ {
		fmt.Printf("%d %c\n", b, b) //这里循坏输出
	}

	array := []byte(str1) //将字符串转换成byte数组，可以理解成为python中的list.
	fmt.Println(array)    //[121 105 110 122 104 101 110 103 106 105 101 119 111 114 108 100]
	array[0] = 72         //也可以写字符'H',在ASCII中他们是一一对应的
	str1 = string(array)
	fmt.Println(str1) //Hinzhengjieworld

	fmt.Println('a' + ('H' - 'h')) //其实就是这些字符在ASCII中的数字相加减，得到的数字为65。
	fmt.Println(0xa)               //表示16进制数字a在ASCII所对应的数字:10
	fmt.Println(0xb)               //表示16进制数字a在ASCII所对应的数字:11
	fmt.Println(0xc)               //表示16进制数字a在ASCII所对应的数字:12

}

// 判断变量的数据类型
func StringC() {

	str1 := "yinzhengjie" + " hello"
	fmt.Println(reflect.TypeOf(str1)) //判断变量“str1”的数据类型:string
	fmt.Println(str1)                 //yinzhengjie hello

	var str2 byte
	str2 = str1[0]                    //让str2=字母"y"在ASCII编码中所对应的数字"121"(我说的说的是十进制的哟)
	fmt.Println(str2)                 //121
	fmt.Println(reflect.TypeOf(str2)) //判断变量"str2"的数据类型 //uint8
}

//用Golang语言实现将小写字母转换成大写字母。
func StringToUpper(s string) (a, b string) {
	if s == "" {
		s = "abcd"
	}
	array := []byte(s) //将字符串转换成字节
	var (
		i   int
		res string
		ss  []string
		sss string
	)
	for i = 0; i < len(s); i++ {
		fmt.Printf(string(array[i] - 32)) //格式化输出，将每一次循环的变量i都保存下来。如果不敲击这一行默认只会保留最后一次循环的值.
		//fmt.Println(string(array[i] - 32))
		res = string(array[i] - 32) //将字节转换成字符串并赋值给res字符串.
		//fmt.Println(res)
		ss = append(ss, res)

	}
	//fmt.Println(res) // 这个只返回最后一个
	sss = strings.Join(ss, "")
	fmt.Println("\n%%---" + sss)
	return res, sss
}

//对比string，byte，rune存取汉字的方式
func StringD() {
	yzj := "尹正杰ABCD"
	yzj2 := []rune(yzj) //rune类型
	yzj3 := []byte(yzj) //byte类型
	for k, v := range yzj {
		fmt.Println(k, "=字符串遍历=>", v)
	}
	fmt.Println()
	for k, v := range yzj2 {
		fmt.Println(k, "=rune遍历=>", v)
	}
	fmt.Println()
	for k, v := range yzj3 {
		fmt.Println(k, "=byte遍历=>", v)
	}
	fmt.Println()
	fmt.Println("让我们来一起看看'rune'的数据类型吧：", reflect.TypeOf(yzj2))
	fmt.Println("让我们来一起看看'byte'的数据类型吧：", reflect.TypeOf(yzj3))
}

// 用rune实现将字符倒着打印即可完成反转功能
func StringRuneE() {
	yzj := "尹正杰 abc 欢 迎你"
	fmt.Println(yzj)
	yzj2 := []rune(yzj) //rune类型
	for i := len(yzj2) - 1; i >= 0; i-- { //从最后一个字符开始倒着打印
		fmt.Printf(string(yzj2[i]))
	}
}

// 使用rune，循环方式变化
func StringRuneF() {
	yzj := "尹 正杰Gl an .g"
	fmt.Println(yzj)
	yzj2 := []rune(yzj) //rune类型
	for k := range yzj2 { //利用数字循环倒着打印，思路跟姿势一一样
		fmt.Printf(string(yzj2[len(yzj2)-k-1]))
	}
}

// 使用rune，循环到中间不操作
func StringRuneG() {
	strs := "北京欢迎你！rggasdfa真的 烦噶风格"
	fmt.Println(strs)
	num := []rune(strs)
	lang := len(num)
	fmt.Println(num)
	for i, j := 0, lang-1; i < j; i, j = i+1, j-1 { //这种思路就是把最后一个字符和第一个字符互换，循环到最中间的那个就不做任何操作
		fmt.Println(i, j)
		num[i], num[j] = num[j], num[i]
	}
	fmt.Println(string(num))
}
