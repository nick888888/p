
package pt

//世界上只有10种人，一种是懂二进制，一种是不懂二进制。
import (
	"math/rand"
	"time"
	"strconv"
	"fmt"
)

var x int64
var s string

//生成随机字符串案例
func PRand() {
	rand.Seed(time.Now().Unix())
	x = rand.Int63()
	s = strconv.FormatInt(x, 36) //表示用36进制来计算x的值。
	fmt.Println(s)               //我们可以在这里对字符串做切片操作，这样就可以随机获取到我们想要到字符了。
}
