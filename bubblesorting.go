package main
// importing fmt package
import ("fmt")
func main(){
 a :=[]int{31, 13, 12, 100}
var i,j int
var n int=4
var temp int 

  for i=1;i<n;i++{
    for j=1;j<n;j++{
      if a[i-1]>a[i]{
        temp=a[i]
        a[i]=a[i-1]
        a[i-1]=temp

    }
      }
  }

fmt.Println("Bubble Sorter")
fmt.Println(a)
}




