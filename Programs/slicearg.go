package main

import "fmt"

func main(){
 slice1 := []int{1,2,3,4}
 fun(slice1...)
}
 
func fun(slice2 ...int){
for _, value := range slice2 {
fmt.Println(value)
}
}