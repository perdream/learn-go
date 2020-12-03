//Go 程（goroutine）是由 Go 运行时管理的轻量级线程。
//go f(x, y, z) 会启动一个新的 Go 程并执行
//f(x, y, z)
package main
import (
    "fmt"
    "time"
)
func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}
func main() {
    go say("world")
    say("hello")
}