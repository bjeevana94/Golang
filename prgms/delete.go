package main
import "fmt"
func main(){
slice1 := []string{"a","b","c","d"}
slice2 := []string{"x","y","z"}
slice1 =append(slice1,slice2...)
fmt.Println(slice1)
slice2 = append(slice1[1:3],slice2[:2]...)
fmt.Println(slice2)
}