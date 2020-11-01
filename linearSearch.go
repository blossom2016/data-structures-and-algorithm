package main

// importing fmt package
import (
	"fmt"
)
func main(){
var arr []int
arr =[]int {10,20,30,40,15, 48, 26, 18, 41, 86, 29, 51, 20};
var key int
key=50
//var flag int
var i int 
for  i=0;i<len(arr);i++ {
        if(key ==arr[i]){
    fmt.Println("Key found")
         
        break
  }
   fmt.Println("Not found") 
  }
  // if flag==1{
  //   fmt.Println("Key not found ")
  // }
}

