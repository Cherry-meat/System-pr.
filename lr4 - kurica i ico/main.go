package main

import(
	"fmt"
	"time"
)

var ch = make(chan string, 1)

type obj struct{

	name string

}

func (o obj) run(){

	for i := 0; i < 10; i++{

		time.Sleep(time.Millisecond)
		fmt.Println("Первым появилось", o.name)

	}

	ch <- fmt.Sprint(o.name)
}

func main(){

	var chicken obj = obj{"Курица"}
	var egg obj = obj{"Яйцо"}

	go egg.run()
	go chicken.run()

	<- ch
	fmt.Println("Спор решен! Первым появилось - ", <- ch)
}