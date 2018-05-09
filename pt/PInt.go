package pt

import (
	"fmt"
	"unsafe"
)

func Pint(){
	var  (
		a int     //包括所有的整数
		b int8    //取值个数2^8,取之范围（-128）-(127)
		c int16
		d int32  //取值个数：2^32，取值范围：(-2^16)-(2^16-1)
		e int64
		x uint     //其中这个"u"，即：unsigned。表示无符号，即只能表示0和所有的正整数。
		y uint8
		z uint16
		p uint32   //取之个数：2^32，取之范围0-2^32-1
		q uint64
	)
	fmt.Println("A=",a,"size=",unsafe.Sizeof(a))
	fmt.Println("B=",b,"size=",unsafe.Sizeof(b))
	fmt.Println("C=",c,"size=",unsafe.Sizeof(c))
	fmt.Println("D=",d,"size=",unsafe.Sizeof(d))
	fmt.Println("E=",e,"size=",unsafe.Sizeof(e))
	fmt.Println("X=",x,"size=",unsafe.Sizeof(x))
	fmt.Println("Y=",y,"size=",unsafe.Sizeof(y))
	fmt.Println("Z=",z,"size=",unsafe.Sizeof(z))
	fmt.Println("P=",p,"size=",unsafe.Sizeof(p))
	fmt.Println("Q=",q,"size=",unsafe.Sizeof(q))

}