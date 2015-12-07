package main
import "fmt"
func main() {
        s := []string{"kala","ashu","tanu","tanusha"}
	argu(s...)  // arguments
}
func argu(s1 ...string) {  //parameters
for _ ,value := range s1 {
fmt.Println(value)
}
}
