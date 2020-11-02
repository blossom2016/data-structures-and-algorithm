package main
// importing fmt package
import ("fmt"
)

func genarrays(i int)[] int{
  arr:= []int{11, 4, 18, 6, 19, 21, 71, 13, 15, 2}

  
  //  rand.Seed(time.Now().UnixNano())
  //  for i=0;i<i;i++{
  //    arr[i]=rand.Intn(999) - rand.Intn(999)
  return arr
   }
  // 
//}
func sortarray(n [] int){
  var i int 
  for i=1;i<len(n);i++{
    j:=i
    for j>0{
  if n[j]<n[j-1]{
    n[j],n[j-1]=n[j-1],n[j]
     
  }
  j = j - 1
  }}

}


func main(){

arrays:=genarrays(5)
fmt.Println("Before Sorting", arrays)
sortarray(arrays)
fmt.Println("After Sorting", arrays)
}




