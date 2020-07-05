package main

import "fmt"
import "sync"

var globalsum = 0.0
var mutex sync.Mutex
var waitGroup sync.WaitGroup

func subsum(arr []float64) {

	for _, val := range arr {
		//mutex.Lock()
		{
			globalsum += val
		}
		//mutex.Unlock()
	}

	waitGroup.Done()

}

func main() {

	const N = 100000
	const M = 10
	waitGroup.Add(M)

	var arr [N]float64

	for i := range arr {
		arr[i] = 1.0
	}

	for i := 0; i < M; i++ {
		go subsum(arr[i*N/M : (i+1)*N/M])
	}

	waitGroup.Wait()

	fmt.Println(globalsum)
}
