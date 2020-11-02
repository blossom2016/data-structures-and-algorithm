package main
// importing fmt package
import ("fmt")
func main(){
var arr []int
arr =[]int {10,20,30,40,15,16};
var key int 
var l int =0
var h int=0
var flag int =0
key=500
for l<h{
  mid :=((l+h)/2)
  if key==arr[mid]{
    flag=1
    break
  }  
  if key>arr[mid]{
    h=mid-1
  }
  else  {
     l=mid+1
  }
  
}
  if flag==1{
    fmt.Println("Key not found ")
  }
}

// else if key >arr[mid]{
//     h=mid-1
//   }
//   else{
//     l=mid+1
//   }