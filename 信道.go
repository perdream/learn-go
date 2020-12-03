// package main
// import "fmt"
// func sum(s []int, c chan int) {
//     sum := 0
//     for _, v := range s {
//         sum += v
//     }
//     c <- sum // 将和送入 c
// }
// func main() {
//     s := []int{7, 2, 8, -9, 4, 0}
//     c := make(chan int)
//     go sum(s[:len(s)/2], c)
//     go sum(s[len(s)/2:], c)
//     x, y := <-c, <-c // 从 c 中接收
//     fmt.Println(x, y, x+y)
// }

// package main

// import (  
//     "fmt"
// )

// func producer(chnl chan int) {  
//     for i := 0; i < 10; i++ {
// 		fmt.Println("写")
//         chnl <- i
//     }
//     close(chnl)
// }
// func main() {  
//     ch := make(chan int)
//     go producer(ch)
//     for {
//         v, ok := <-ch
//         if ok == false {
//             break
//         }
//         fmt.Println("Received ", v, ok)
//     }
// }

// package main

// import (  
//     "fmt"
// )

// func producer(chnl chan int) {  
//     for i := 0; i < 10; i++ {
//         chnl <- i
//     }
//     close(chnl)
// }
// func main() {  
//     ch := make(chan int)
//     go producer(ch)
//     for v := range ch {
//         fmt.Println("Received ",v)
//     }
// }

package main

import (  
    "fmt"
)

func digits(number int, dchnl chan int) {  
    for number != 0 {
        digit := number % 10
        dchnl <- digit
        number /= 10
    }
    close(dchnl)
}
func calcSquares(number int, squareop chan int) {  
    sum := 0
    dch := make(chan int)
    go digits(number, dch)
    for digit := range dch {
        sum += digit * digit
    }
    squareop <- sum
}

func calcCubes(number int, cubeop chan int) {  
    sum := 0
    dch := make(chan int)
    go digits(number, dch)
    for digit := range dch {
        sum += digit * digit * digit
    }
    cubeop <- sum
}

func main() {  
    number := 589
    sqrch := make(chan int)
    cubech := make(chan int)
    go calcSquares(number, sqrch)
    go calcCubes(number, cubech)
    squares, cubes := <-sqrch, <-cubech
    fmt.Println("Final output", squares+cubes)
}
