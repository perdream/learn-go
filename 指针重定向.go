package main
import "fmt"
type Vertex struct {
    X, Y float64
}
//接收者无论是指针类型还是值类型都可以接收指针类型或者值类型(例如：*Vertex和Vertex 都可以接收Vertex对象或者&Vertex)，只不过只有指针类型接收者可以修改指向的接收者字段值
func (v *Vertex) Scale(f float64) {
    //可修改
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v Vertex) Scale2(f float64){

	fmt.Println(f)

	//只能使用字段，不可修改,以下操作不起作用
    v.X = v.X * f * 10
    v.Y = v.Y * f * 10
}

func ScaleFunc(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    // v := Vertex{3, 4}
	// v.Scale(2)
	
	// ScaleFunc(&v, 10)
	
	// //指针为接收者的方法被调用时，接收者既能为值又能为指针
    // p := &Vertex{4, 3}
    // p.Scale(3)
    // ScaleFunc(p, 8)
	// fmt.Println(v, p)
	
	// v := Vertex{3, 4}
	// v.Scale2(2)
	// fmt.Println(v)
	// ScaleFunc(&v, 10)
	
	//指针为接收者的方法被调用时，接收者既能为值又能为指针
    p := &Vertex{4, 3}
    p.Scale2(3)
    fmt.Println(p)
}