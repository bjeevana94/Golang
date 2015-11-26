package main
import "fmt"
func main(){
var s []string
[]string = {"kala","ashu"}
variad("hello:",s...);
}
func variad(my ...s){
for _,value:=range my{
fmt.Println(value)
}
}