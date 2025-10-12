package main

import(
	"fmt"
	"time"
)

// ch - канал для передачи сообщений о завершении забега
var ch = make(chan string)

type animal struct{
	name string  
	speed int    
}

// run - метод, имитирующий бег животного
func (an animal) run (){
	// Цикл бега на 50 метров
	for i := 0; i < 50; i++ {
		time.Sleep(time.Duration(an.speed) * time.Millisecond)
		fmt.Println("Животное", an.name ,"пробежал", i, " метров.")
	}

	// Отправка имени животного в канал по завершении забега
	ch <- an.name
}

func main(){
	var rabbit animal = animal{"Заяц", 10}
	var turtle animal = animal{"Черепаха", 10}

	go rabbit.run() 
	go turtle.run()
	
	// Ожидание и вывод имени первого животного, завершившего забег
	fmt.Println(<- ch)

}
