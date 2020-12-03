package main
import "fmt"

//类型断言 提供了访问接口值底层具体值的方式。
func main() {
    var i interface{} = "hello"
	s := i.(string)
	//hello
    fmt.Println(s)
	s, ok := i.(string)
	//hello true
    fmt.Println(s, ok)
	f, ok := i.(float64)
	//f 被赋予类型的零值，ok返回false
    fmt.Println(f, ok)
    f = i.(float64) // 报错(panic)
    fmt.Println(f)
}