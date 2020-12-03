//信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道: ch := make(chan int, 100)
//仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。

//例子1
// package main
// import "fmt"
// func main() {
//     ch := make(chan int, 2)
//     ch <- 1
//     fmt.Println(<-ch)
//     fmt.Println(<-ch)
// }

//例子2
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	//创建一个容量为3的缓冲信道
// 	ch := make(chan string, 3)
// 	ch <- "naveen"
// 	ch <- "paul"
// 	fmt.Println("capacity is", cap(ch))   //capacity is 3
// 	fmt.Println("length is", len(ch))     //length is 2
// 	fmt.Println("read value", <-ch)       //read value naveen
// 	fmt.Println("new length is", len(ch)) //new length is 1
// }


//例子3
// package main
// import (
// 	"fmt"
// 	"time"
// )

// func write(ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 		fmt.Println("successfully wrote", i, "to ch")
// 	}
// 	close(ch)
// }

// func main(){
// 	ch := make(chan int, 2)
// 	go write(ch)
// 	time.Sleep(2 * time.Second)
// 	for v := range ch {
// 		fmt.Println("read value", v,"from ch")
// 		time.Sleep(2 * time.Second)
// 	}
// }

//例子4 死锁
// package main
// import (
// 	"fmt"
// )
// func main() {
// 	ch := make(chan string, 2)
// 	ch <- "naveen"
// 	ch <- "paul"
// 	ch <- "steve" //此处报死锁错误
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

//例子5
package main

import (
    "fmt"
    "sync"
    "time"
)

func process(i int, wg *sync.WaitGroup) {
	//Done方法减少WaitGroup计数器的值-1，应在线程的最后执行。
	defer wg.Done()
    fmt.Println("started Goroutine ", i)
    time.Sleep(2 * time.Second)
    fmt.Printf("Goroutine %d ended\n", i)
}

/*
    WaitGroup用于等待一组线程的结束。
    父线程调用Add方法来设定应等待的线程的数量。
    每个被等待的线程在结束时应调用Done方法。
    同时，主线程里可以调用Wait方法阻塞至所有线程结束。
*/
func main() {
    no := 3
    var wg sync.WaitGroup
    //并发协程
    for i := 0; i < no; i++ {
        /*
            Add方法向内部计数加上delta，delta可以是负数；
            如果内部计数器变为0，Wait方法阻塞等待的所有线程都会释放，
            如果计数器小于0，方法panic。
        */
        wg.Add(1)
        go process(i, &wg) //开启并发协程
    }
    //Wait方法阻塞直到WaitGroup计数器减为0。
    wg.Wait() //在此阻塞
    fmt.Println("over")
}