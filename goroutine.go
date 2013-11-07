package main

func main() {
    myChan := make(chan int)
    foo(myChan)
    fooIn(myChan)
    fooOut(myChan)
}

func foo(c chan int) {
    c <- 0
}

func fooIn(c chan<- int) {
    c <- 0
    // invalid operation: <-c (receive from send-only type chan<- int)
    // <- c
}

func fooOut(c <-chan int) {
    <- c
    // invalid operation: c <- 0 (send to receive-only type <-chan int)
    // c <- 0
}