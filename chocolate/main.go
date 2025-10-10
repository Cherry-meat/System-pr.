package main

import "fmt"

func main() {
	var money, price, wrap int

	fmt.Print("Введите количество денег: ")
	fmt.Scan(&money)
	fmt.Print("Введите цену за 1 шоколадку: ")
	fmt.Scan(&price)
	fmt.Print("Введите количество оберток для бесплатной шоколадки: ")
	fmt.Scan(&wrap)

	if money < 0 || price <= 0 || wrap <= 0 {
		fmt.Println("Ошибка: деньги, цена и количество оберток должны быть больше 0")
		return
	}

	total := recursiveChocolate(money, price, wrap, 0)
	fmt.Printf("Всего можно получить шоколадок: %d\n", total)
}

func recursiveChocolate(money, price, wrap, wrappers int) int {
	if money < price && wrappers < wrap {
		return 0
	}

	chocolatesFromMoney := 0
	chocolatesFromWrappers := 0

	if money >= price {
		chocolatesFromMoney = money / price
		money %= price
		wrappers += chocolatesFromMoney
		fmt.Printf("Куплено %d шоколадок. Осталось денег: %d, оберток: %d\n", 
			chocolatesFromMoney, money, wrappers)
	}

	if wrappers >= wrap {
		chocolatesFromWrappers = wrappers / wrap
		remainingWrappers := wrappers % wrap
		wrappers = remainingWrappers + chocolatesFromWrappers
		fmt.Printf("Обменяли обертки на %d шоколадок. Теперь оберток: %d\n", 
			chocolatesFromWrappers, wrappers)
	}

	return chocolatesFromMoney + chocolatesFromWrappers + 
		recursiveChocolate(money, price, wrap, wrappers)
}