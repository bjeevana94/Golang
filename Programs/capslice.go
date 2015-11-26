package main
import "fmt"
func main(){
capslice := make([]string , 5 ,10)
capslice[0]="kala"
capslice[1]="jeevana"
capslice[2]="tanu"
capslice[3]="ashu"
capslice[4]="tanusha"
capslice[5]="family"
for _,value := range capslice{
fmt.Println(value)
}
}