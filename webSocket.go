// http://talks.golang.org/2012/chat.slide
package main
import (
    "io"
    "log"
    "net/http"
    "html/template"

    "code.google.com/p/go.net/websocket"
)

const listenAddr = "localhost:4000"

func main() {
    http.HandleFunc("/", rootHandler)
    http.Handle("/socket", websocket.Handler(socketHandler))
    err := http.ListenAndServe(listenAddr, nil)
    if err != nil {
        log.Fatal(err)
    }
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    rootTemplate.Execute(w, listenAddr)
}

var rootTemplate = template.Must(template.New("root").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<script>
    websocket = new WebSocket("ws://{{.}}/socket");
    websocket.onmessage = onMessage;
    websocket.onclose = onClose;
</script>
</html>
`))


type socket struct {
    io.ReadWriter
    done chan bool
}

func (s socket) Read(b []byte) (int, error)  { return s.conn.Read(b) }
func (s socket) Write(b []byte) (int, error) { return s.conn.Write(b) }

func (s socket) Close() error {
    s.done <- true
    return nil
}

func socketHandler(ws *websocket.Conn) {
    s := socket{ws, make(chan bool)}
    go match(s)
    <-s.done
}

func match(value bool) {
    return true
}