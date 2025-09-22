package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	n := 1000
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i] = rand.Intn(10001) 
	}

	max14 := -1   
	max7 := -1    
	max2 := -1    
	max := -1     

	for _, num := range sequence {
		if num%14 == 0 {
			if num > max14 {
				max14 = num
			}
		} else if num%7 == 0 && num&14 != 0 {
			if num > max7 {
				max7 = num
			}
		} else if num%2 == 0 && num&14 != 0 {
			if num > max2 {
				max2 = num
			}
		}
		
		if num > max {
			max = num
		}
	}
	if max * max14 > max2 * max7{
		fmt.Printf("%d * %d = %d", max,max14,max*max14)
	}else{
		fmt.Printf("%d * %d = %d", max2,max7,max2*max7)
	}
}