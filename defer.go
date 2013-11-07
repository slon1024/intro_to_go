package main
import "fmt"

func main() {
	fmt.Println(c())
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func c() (i int) {
	defer func() { i += 10 }()
	return 1
}