package main
import (
    "fmt"
    "math"
)

//接口定义
type Abser interface {
    Abs() float64
}
func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}
    a = f  // a MyFloat 实现了 Abser
    a = &v // a *Vertex 实现了 Abser
	// 下面一行，v 是一个 Vertex（而不是 *Vertex),接收者类型要严格匹配，不像是给结构体添加方法(不论接收者是引用还是值)
    // 所以没有实现 Abser。
    a = v
    fmt.Println(a.Abs())
}
type MyFloat float64
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}
type Vertex struct {
    X, Y float64
}
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}