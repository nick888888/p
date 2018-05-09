// 临时记录，未分类
package pt

import (
	"fmt"
	"strconv"
	"os"
	"log"
	"bufio"
	"io"
	"io/ioutil"
)

// 循环
func PA_a() {

	fmt.Println("遍历1")
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()
	//0
	//1
	//2
	//3
	//4
	//5

	fmt.Println("遍历2")
	i := 1
	for i < 10 {
		fmt.Println(i)
		i = i + 2
	}
	fmt.Println()
	//1
	//3
	//5
	//7
	//9

	fmt.Println("遍历3")
	j := 8
	for {
		j = j + 1
		fmt.Println(j)
		if j > 10 {
			break
		}
	}
	fmt.Println()
	//9
	//10
	//11

	fmt.Println("遍历4")
	s := "尹正杰2017"
	for i, arg := range s {
		fmt.Printf("%d %c\n", i, arg)
	}
	//0 尹
	//3 正
	//6 杰
	//9 2
	//10 0
	//11 1
	//12 7
}

//strconv
func PA_b() {
	s1 := "2017"
	n1, err := strconv.Atoi(s1) //用"n"和"err"两个参数来接受“strconv.Atoi”的值，如果正常就把值传给变量"n",如果异常就把值传给"err"。
	if err != nil { //如果错误信息不为空，那么就执行大括号的代码。
		fmt.Println("'strconv.Atoi'传入的参数有问题哟！看报错内容是>>>:", err) //将接受的错误信息打印出来。
		return                                                 //出现这行代码就不会往下走啦，直接就终止这个函数了。
	}
	fmt.Println(n1)
	fmt.Println("1111")

	s2 := "尹正杰2017"
	n2, err := strconv.Atoi(s2) //用"n"和"err"两个参数来接受“strconv.Atoi”的值，如果正常就把值传给变量"n",如果异常就把值传给"err"。
	if err != nil { //如果错误信息不为空，那么就执行大括号的代码。
		fmt.Println("'strconv.Atoi'传入的参数有问题哟！看报错内容是>>>:", err) //将接受的错误信息打印出来。
		return                                                 //出现这行代码就不会往下走啦，直接就终止这个函数了。
	}
	fmt.Println(n2)
	fmt.Println("2222")
}

// os 接收输入并计算
func PA_c() {
	var (
		x int
		y int
		//res int
	)

	fmt.Println("下面是打印命令行的输入")
	for _, v := range os.Args {
		fmt.Println(v)
	}
	fmt.Println()

	fmt.Println("下面是对输入的值进行计算")
	x, _ = strconv.Atoi(os.Args[1]) //将字符串转换成数字
	y, _ = strconv.Atoi(os.Args[3])
	switch os.Args[2] {
	case "+":
		fmt.Printf("运算结果是：%d\n", x+y)
	case "-":
		fmt.Printf("运算结果是：%d\n", x-y)
	case "*":
		fmt.Printf("运算结果是：%d\n", x*y)
	case "/":
		fmt.Printf("运算结果是：%d\n", x/y)
	case "%":
		fmt.Printf("运算结果是：%d\n", x%y)
	}
	//fmt.Println(os.Args[1])

}

// 文件的写入，如果以前有这个文件，会被清空再写入
func PA_d() {
	f, err := os.Create("a.txt") //姿势一：打开一个文件，如果没有就创建，如果有这个文件就清空文件内容,需要用两个变量接受相应的参数
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("yinzhengjie\n") //往文件写入相应的字符串。
	f.Close()
}

// 以追加的方式打开一个文件（相当于python中的"a"）;  //如果以前有这个文件，会在文件后面增加字符
func PA_e() {
	f, err := os.OpenFile("a.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644) //表示最佳的方式打开文件，如果不存在就创建，打开的模式是可读可写，权限是644
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("add\n")
	f.Close()
}

// 修改文件内容-随机写入（自定义插入的位置，相当python重的seek方法）
func PA_f() {
	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("abcdefghijklmn\n")
	f.Seek(1, os.SEEK_SET) //表示文件的其实位置，从第二个字符往后写入。
	f.WriteString("$$$")
	f.Close()
	//a$$$efghijklmn  这个是替换中间的字符，不是写入
}

// 顺序读取文件内容
//姿势一：读取文件内容
func PA_Read_a() {
	//f,err := os.OpenFile("a.txt",os.O_CREATE|os.O_RDWR,0644) //打开该文件，但是不建议用这行命令，用下面的专门用来读取文件的。
	f, err := os.Open("a.txt") //打开一个文件，如果这个文件不存在的话就会报错。
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f) //取出文件的内容
	for {
		line, err := r.ReadString('\n') //表示以"\n"为分隔符，遇到"\n"就换行。
		fmt.Print(line)
		if err == io.EOF { //读到文件结尾后就终止文件。
			break
		}
	}
	f.Close()
}

// 随机读取文件内容
func PA_Read_b() {
	f, err := os.OpenFile("yinzhengjie.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n') //表示以"\n"为分隔符，遇到"\n"就换行。
		fmt.Print(line)
		if err == io.EOF { //读到文件结尾后就终止文件。因此如果你这么干了的话，你始终只能读取到第一行的内容
			break
		}
		//fmt.Print(line)
	}
	f.Close()
}

//用ioutil模块读取文件内容；
func PA_Read_c() {
	buf, err := ioutil.ReadFile("D:\\Golang环境\\Golang Program\\Golang lesson\\Day4\\单词.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}

//用Golang编写99乘法表；生成乘法表
func PA_jj() {
	var top int = 10
	//f,err := os.Create("乘法表.txt")
	f, err := os.OpenFile("乘法表.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < top; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fprintf(f, "%d * %d = %2d\t", i, j, i*j)
		}
		fmt.Fprintln(f)
	}
	f.Close()
}

// 用Golang将九九乘法表读取出来
func PA_jj_read(){
	f,err := os.Open("乘法表.txt") //打开一个文件，如果这个文件不存在的话就会报错。
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f) //取出文件的内容
	for {
		line,err := r.ReadString('\n') //表示以"\n"为分隔符，遇到"\n"就换行。
		if err == io.EOF{ //读到文件结尾后就终止文件。
			break
		}
		fmt.Print(line)
	}
	f.Close()
}
//1 * 1 =  1
//2 * 1 =  2	2 * 2 =  4
//3 * 1 =  3	3 * 2 =  6	3 * 3 =  9
//4 * 1 =  4	4 * 2 =  8	4 * 3 = 12	4 * 4 = 16
//5 * 1 =  5	5 * 2 = 10	5 * 3 = 15	5 * 4 = 20	5 * 5 = 25
//6 * 1 =  6	6 * 2 = 12	6 * 3 = 18	6 * 4 = 24	6 * 5 = 30	6 * 6 = 36
//7 * 1 =  7	7 * 2 = 14	7 * 3 = 21	7 * 4 = 28	7 * 5 = 35	7 * 6 = 42	7 * 7 = 49
//8 * 1 =  8	8 * 2 = 16	8 * 3 = 24	8 * 4 = 32	8 * 5 = 40	8 * 6 = 48	8 * 7 = 56	8 * 8 = 64
//9 * 1 =  9	9 * 2 = 18	9 * 3 = 27	9 * 4 = 36	9 * 5 = 45	9 * 6 = 54	9 * 7 = 63	9 * 8 = 72	9 * 9 = 81

