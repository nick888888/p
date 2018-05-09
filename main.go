package main

import (
	"fmt"
	"strings"
	"time"
	"lvmama/intlflight_protocol/prototype"
	"math"
	//"lvmama/p/pt"
	"lvmama/p/pt"
)

func main() {
	//fmt.Println("main start")

	//tString()
	//testTime()
	//testAppend()
	//testIfElse()
	//testStructCopy()
	//testInt() //测试取整




	pt.PA_c()  //执行 go run main.go 2 + 6

	//fmt.Println("main end")
}

func tString() {
	fmt.Println(strings.Count("five", ""))  //5
	fmt.Println(strings.Count("fi we", "")) //6
	fmt.Println(strings.Count("", ""))      //1

	//fmt.Println(strings.Contains("seafood", "foo")) //true
	//fmt.Println(strings.Contains("seafood", ""))   //true //这里要特别注意
	//fmt.Println(strings.Contains("", ""))           //true //这里要特别注意
	fmt.Println(strings.ContainsAny("failure", "u & i"))  // true
	fmt.Println(strings.ContainsAny("failure", "us  is")) // true
	//fmt.Println(strings.ContainsAny("foo", ""))          // false
	//fmt.Println(strings.ContainsAny("", ""))             // false
	//fmt.Println(strings.ContainsRune("我是中国", '我'))  //true
	//fmt.Println(strings.EqualFold("Go", "go"))      // ture 大小写忽略
	//fmt.Println(strings.ContainsAny("foo", ""))     // false

}

func testTime() {
	departureDateReq, err := time.Parse("2006-01-02", "2018-05-29")

	if err != nil {
		fmt.Println(err.Error() + "-----err")
	}
	fmt.Println("start*****")
	fmt.Println(departureDateReq.Year())  //2018
	fmt.Println(departureDateReq.Month()) //May
	fmt.Println(departureDateReq.Day())   //29

	println(departureDateReq.Year())  //2018 //go build main.go的时候，不是按照顺序输出
	println(departureDateReq.Month()) //5    //go build main.go的时候，不是按照顺序输出
	println(departureDateReq.Day())   //29   //go build main.go的时候，不是按照顺序输出
	fmt.Println("end----")
}

type ren struct {
	name string
	code int
}

func testAppend() {
	ss := []*ren{}
	a := ren{
		name: "ali",
		code: 1,
	}
	b := ren{
		name: "bom",
		code: 2,
	}
	ss = append(ss, &a)
	ss = append(ss, &b)

	for k, v := range ss {
		fmt.Printf("k:%d\n", k) // for range 的 k 是从0开始。
		fmt.Println(v.name)
		ss = append(ss, &ren{// 这里追加不会影响到正在循环的ss循环次数，最终数据会追加到ss中
			name: "add",
			code: k + 5,
		})
		//fmt.Println("当前的k是:%d", k)  //这里输出是  "当前的k是:%d 0"
		fmt.Printf("----当前的k是:%d\n", k) //这里输出是  "当前的k是:0"
	}

	fmt.Println("222--for range")
	for _, v := range ss {
		fmt.Printf("--88--code:%v----name:%v\n", v.code, v.name)
	}

}

func testIfElse() {
	a := 2
	b := 3
	if a > 0 { // 同时满足只走a>0
		fmt.Println("a>0")
	} else if b > 0 {
		fmt.Println("b>0")
	}
	fmt.Println("end")

}

func testStructCopy() (prices []*prototype.Price) {

	if len(prices) > 0 {
		fmt.Printf("len(prices):%d\n", len(prices))
	}
	if len(prices) <= 0 {
		fmt.Printf("len(prices):%d\n", len(prices))
	}

	for k, _ := range prices { //这里prices的len=0，并不影响后面ffff输出
		fmt.Printf("%d", k)
		fmt.Printf("rrrr")
	}
	fmt.Printf("-----ffff\n")

	cabinType1 := &prototype.CabinType{
		Index: 1,
		Data:  22,
	}
	cabinType2 := &prototype.CabinType{
		Index: 1,
		Data:  23,
	}

	cabinTypes := []*prototype.CabinType{}

	cabinTypes = append(cabinTypes, cabinType1, cabinType2)

	price := &prototype.Price{
		PaxType:    "Adult",
		CabinTypes: cabinTypes,
		FareType:   prototype.FareTypeChange,
	}

	prices = append(prices, price)
	// 如果fareType不等于1时，需要复制一份价格FareType设置为1
	if price.FareType != prototype.FareTypeNormal {
		fmt.Printf("----%d\n", price.FareType)
		copyPrice := *price
		copyPrice.FareType = prototype.FareTypeNormal
		prices = append(prices, &copyPrice)
	}

	for k, v := range prices {
		fmt.Printf("**  k:%v,price:%v\n", k, v.FareType) //复制成功
	}
	return
}

func testInt() {
	a := 8
	b := 3
	c := a / b
	fmt.Printf("aaa--%d\n", c)  //这里正常输出  aaa--2
	fmt.Println("bbb--%d\n", c) // 这里输出  bbb--%d  2

	//fa := 15
	//fb := 4
	d := math.Floor(99.556) //向下取整99
	fmt.Println(d)

	e := math.Ceil(99.556) //向上取整100
	fmt.Println(e)

}

// 四舍五入 没验证http://sugarmanman.blog.163.com/blog/static/810790802013812104859519/
func round(val float64, places int) float64 {
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}
