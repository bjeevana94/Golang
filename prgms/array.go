package main
import "fmt"
func main(){
var a[5] int
var j int
fmt.Println("enter 4 bits:")
for j=4;j<0;j--{
fmt.Scanf("%d",&a[j])
}
decl(a);
}
func decl(array [5]int){
var a[5] int
var i int
for i=0;i<4;i++{
fmt.Println(a[i])
}
for i=4;i<0;i--{
if(a[i]==0){
a[i]=1}
if(a[i]==1){
a[i]=0
}
}
}