package main
import "fmt"
func main() {
	//i 为空接口类型，i可保存任何类型的值
    var i interface{}
    describe(i)
    i = 42
    describe(i)
    i = "hello"
    describe(i)
}

//空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数
func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}