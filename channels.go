//http://talks.golang.org/2012/waza.slide
package main
import "time"

func main() {
    deltaT := 1
    
    timerChan := make(chan time.Time)
    go func() {
        time.Sleep(deltaT)
        timerChan <- time.Now() // send time on timerChan
    }()
    // Do something else; when ready, receive.
    // Receive will block until timerChan delivers.
    // Value sent is other goroutine's completion time.
    completedAt := <-timerChan
}