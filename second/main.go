package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	// Создаем случайную последовательность из 1000 элементов
	n := 1000
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i] = rand.Intn(10001) // числа от 0 до 10000
	}

	// Инициализация переменных для хранения максимальных чисел
	max14 := -1   // максимальное число, кратное 14
	max7 := -1    // максимальное число, кратное 7 но не 2
	max2 := -1    // максимальное число, кратное 2 но не 7
	max := -1     // просто максимальное число

	// Проходим по всем элементам последовательности
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
		
		if num > max {
			max = num
		}
	}

	result := -1

	// Проверяем возможные комбинации для получения произведения, кратного 14
	if max14 != -1 && max != -1 {
		candidate := max14 * max
		if candidate > result {
			result = candidate
		}
	}

	if max7 != -1 && max2 != -1 {
		candidate := max7 * max2
		if candidate > result {
			result = candidate
		}
	}

	// Выводим результат
	fmt.Printf("\nМаксимальное число кратное 14: %d\n", max14)
	fmt.Printf("Максимальное число кратное 7 но не 2: %d\n", max7)
	fmt.Printf("Максимальное число кратное 2 но не 7: %d\n", max2)
	fmt.Printf("Абсолютный максимум: %d\n", max)
	
	fmt.Printf("\nРезультат R: %d\n", result)
	
	if result == -1 {
		fmt.Println("Подходящего числа не найдено")
	} else {
		fmt.Printf("Проверка (показывается остаток): %d %% 14 = %d\n", result, result%14)
	}
}