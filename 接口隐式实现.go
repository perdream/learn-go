package main
import "fmt"
type I interface {
    M()
}
type T struct {
    S string
}
// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。(只要是实现了接口I的所有方法，则表示类型T继承了接口I)
func (t T) M() {
    fmt.Println(t.S)
}
func main() {
    var i I = T{"hello"}
    i.M()
}