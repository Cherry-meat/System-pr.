package main

import(
	"fmt"
	"time"
)

var ch = make(chan string)

type animal struct{
	name string
	speed int
}

func (an animal) run (){

	for i := 0; i < 50; i++ {

		time.Sleep(time.Duration(an.speed) * time.Millisecond)
		fmt.Println("Животное", an.name ,"пробежал", i, " метров.")

	}

	ch <- an.name
}

func main(){

	var rabbit animal = animal{"Заяц", 10}
	var turtle animal = animal{"Черепаха", 10}

	go rabbit.run() 
	go turtle.run()
	
	fmt.Println(<- ch)
}