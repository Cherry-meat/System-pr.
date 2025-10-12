package main

import(
	"fmt"
	"time"
)

// ch - буферизированный канал для передачи сообщений о завершении, позволяет записать одно сообщение без блокировки
var ch = make(chan string, 1)

type obj struct{
	name string 
}

// run - метод, имитирующий процесс "появления" объекта
func (o obj) run(){
	for i := 0; i < 10; i++{
		time.Sleep(time.Millisecond)
		fmt.Println("Первым появилось", o.name)
	}

	// Отправка имени объекта в канал по завершении процесса
	ch <- fmt.Sprint(o.name)
}

func main(){
	var chicken obj = obj{"Курица"}
	var egg obj = obj{"Яйцо"}

	go egg.run()
	go chicken.run()

	// Ожидание первого сообщения из канала (но не используем его)
	<- ch
	// Чтение второго сообщения из канала и вывод результата
	fmt.Println("Спор решен! Первым появилось - ", <- ch)

}
