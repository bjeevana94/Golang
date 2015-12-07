package main
import "fmt"
func main(){
kala := make(map[int] string)
kala[1]="hai"
kala[2]="ashu"
kala[3]="tanu"
for _,value :=range kala{
fmt.Println(value)
}
}