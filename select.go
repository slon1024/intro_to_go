// http://talks.golang.org/2012/waza.slide
package main
import "fmt"

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    select {
        case v := <-ch1:
            fmt.Println("channel 1 sends", v)
        case v := <-ch2:
            fmt.Println("channel 2 sends", v)
        default: // optional
            fmt.Println("neither channel was ready")
    }
}