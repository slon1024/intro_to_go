// http://talks.golang.org/2012/chat.slide
package main
import (
    "fmt"
    "log"
    "net/http"

    "code.google.com/p/go.net/websocket"
)

const listenAddr = "localhost:4000"

func main() {
    http.Handle("/", websocket.Handler(handler))
    err := http.ListenAndServe(listenAddr, nil)
    if err != nil {
        log.Fatal(err)
    }
}

func handler(c *websocket.Conn) {
    var s string
    fmt.Fscan(c, &s)
    fmt.Println("Received:", s)
    fmt.Fprint(c, "How do you do?")
}