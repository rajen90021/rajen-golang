// //////wg.add()-> used to set the number og goroutines to wait for
// ///wg.wait()-> wait is used to block the execution until all gorutine have finised
// /wg.done()-> each of the goroutine runs and call done when fineshed
package main

import (
	"fmt"
	"sync"
	"time"
)

// var data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// func fetchdata(i int) {
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(data[i])
// }
// func main() {
// 	//current := time.Now()
// 	for i := 0; i < 10; i++ {
// 		go fetchdata(i)
// 	}
// 	time.Sleep(2 * time.Second)/// it will stop the currrent main thread and give a chance t0 oexecuted to goroutine
// 	fmt.Println("ekekek")

// }

var data = []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10}

func fetdata(i int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println(data[i])
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10) ///saying 10 goroutine to be executed
	for i := 0; i < 10; i++ {
		go fetdata(i, &wg) //passing waitgroup to  the fetdata
	}
	wg.Wait() ///wait is used to block the execution until all gorutine have finised
	fmt.Println("main thread executed lastly")
}
