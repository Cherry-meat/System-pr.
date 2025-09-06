package main

import "fmt"
import "math"

func main(){
  firstWay(5,2)
  secondWay(4, 52)
  thirdWay(3, 12)
  fourthWay(7, 2)
  fifthWay(3,2)
  fmt.Println(sixthWay(7,15))
}

func firstWay(a int, b int){
  result := 0
  for i := 0; i  <b; i++{
    result += a
  }
  fmt.Println(result)
}
func secondWay(a float32, b float32){
  fmt.Println(int( a/(1/b)))
}
func thirdWay(a float64, b float64){
  fmt.Println(int(math.Pow(2, math.Log2(a)+ math.Log2(b))))
}
func fourthWay(a float64, b float64){
  fmt.Println(math.Round(1 - (a+b)/(math.Tan((math.Atan(a)+ math.Atan(b))))))
}

func fifthWay(a int, b int){
    if a == 0 || b == 0 {
        fmt.Println("0")
    }
    result := 0
    for b > 0 {
        if b&1 == 1 {
            result += a
        }
        a <<= 1
        b >>= 1
    }
  fmt.Println(result)
}

func sixthWay(a int,b int) int{
  if (a == 0 || b == 0){
    return 0;
  } 
    if (b == 1){
    return a;
  } 
    if (b == -1){
     return -a;
  }

  
  halfB := b / 2
    half := sixthWay(a, halfB)
    remainder := b % 2
    if remainder != 0 {
        if remainder == 1 || remainder == -1 {
            return half + half + a
        }
    }
    return half + half
}