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

	if money <= 0 || price <= 0 || wrap <= 0 { // Ну если человек захочет в долг покупать или чтобы ему доплатили за шоколадки
		fmt.Println("Ошибка: деньги, цена и количество оберток должны быть больше 0")
		return
	}

	total := raschetChoco(money, price, wrap, 0) // Чтобы узнать сколько всего шоколадок можно получить
	fmt.Printf("Всего можно получить шоколадок: %d\n", total)
}

func raschetChoco(money, price, wrap, wrappers int) int { //Узнаем сколько шоколадок можно купить

	if money < price {
		return 0
	}

	chocoMoney := 0

	if money >= price {
		chocoMoney = money / price
		money %= price
		wrappers += chocoMoney
		fmt.Printf("Куплено %d шоколадок. Осталось денег: %d, оберток: %d\n", chocoMoney, money, wrappers)
	}

	return chocoMoney + WrapChoco(wrap, wrappers)
}

func WrapChoco(wrap, wrappers int) int { //Находятся шоколадки за обертки, в том числе шоколадки полученные за обертки тоже учитываются.

	chocoWrap := 0
	chwr := 0

	for true { // будет выполняться чтобы найти все шоколадки, за все обертки 
		if wrappers >= wrap {
			chocoWrap = wrappers / wrap
			remWrappers := wrappers % wrap
			wrappers = remWrappers + chocoWrap
		} else {
			break
		}
		chwr += chocoWrap
	}

	fmt.Printf("Обменяли обертки на %d шоколадок. Теперь оберток: %d\n", chwr, wrappers)
	return chwr
}
