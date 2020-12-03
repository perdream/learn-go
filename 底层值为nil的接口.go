package main
import "fmt"
type I interface {
    M()
}
type T struct {
    S string
}
//接收者可以是实现了接口的类型的nil值
func (t *T) M() {
    if t == nil {
        fmt.Println("")
        return
    }
    fmt.Println(t.S)
}
func main() {
	var i I
	describe(i)
	var t *T
	describe(t)
    i = t
    describe(i)
    i.M()
    i = &T{"hello"}
    describe(i)
    i.M()
}
func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}