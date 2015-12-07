package main
import "fmt"
func main() {
        s := []int{1,2,3,4}
	argu(s...)
}
func argu(s1 ...int) {
fmt.Println(s1)
}
