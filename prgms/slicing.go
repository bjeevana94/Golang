package main
import "fmt"
func main(){
var s string
fmt.Println("enter string")
fmt.Scan(&s)
fmt.Println(s[:1])
fmt.Println(s[:5])
fmt.Println(s[1:3])
fmt.Println(s[:])
}