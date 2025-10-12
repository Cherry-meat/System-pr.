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
	// Заполняем массив случайными числами от 0 до 10000
	for i := 0; i < n; i++ {
		sequence[i] = rand.Intn(10001) 
	}

	max14 := -1   // максимальное число, кратное 14
	max7 := -1    // максимальное число, кратное 7 (но не 14)
	max2 := -1    // максимальное число, кратное 2 (но не 14)
	max := -1     // максимальное число в последовательности

	// Проходим по всем числам последовательности
	for _, num := range sequence {
		if num%14 == 0 {
			if num > max14 {
				max14 = num
			}
		} else if num%7 == 0 {
			if num > max7 {
				max7 = num
			}
		} else if num%2 == 0 {
			if num > max2 {
				max2 = num
			}
		}
		
		// Обновляем абсолютный максимум
		if num > max {
			max = num
		}
	}

	// Выбираем максимальное произведение
	if max * max14 > max2 * max7 {
		fmt.Printf("%d * %d = %d", max, max14, max*max14)
	} else {
		fmt.Printf("%d * %d = %d", max2, max7, max2*max7)
	}
}
