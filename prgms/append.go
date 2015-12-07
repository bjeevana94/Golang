package main
import "fmt"
func main(){
slice1 := []string{"happines" , "love", "Pride", "money"}
slice2 := []string{"1","2","3","4"}
slice2 = append(slice2,slice1...)
fmt.Println(slice2)
}