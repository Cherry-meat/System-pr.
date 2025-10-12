package main
import (
	"fmt"
	"sync"
)
// создается 10 горутинок, которые одновременно работают. "Sync" нужен, чтобы дождаться выполнения всех горутин.
var wg sync.WaitGroup

func main(){
	
	wg.Add(10)
	fmt.Println("Start")
	for i:=0; i<10; i++{
		go sum(i)
	}
	wg.Wait()
	fmt.Println("End")
}

func sum(i int){
	fmt.Println(i)
	wg.Done()

}
