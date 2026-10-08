package main

import (
	"fmt"
	"sync"
)

// func worker(id int, jobs <-chan int, hasil chan<- string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for job := range jobs {
// 		fmt.Printf("Worker %d sedang memproses job #%d\n", id, job)

// 		hasil <- fmt.Sprintf("Hasil job#%d dari Worker %d", job, id)
// 	}
// }

func worker(jobs <-chan int, hasil chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		hasil <- fmt.Sprintf("job %d", job)
	}
}

func main() {

	jobs := make(chan int, 2)
	hasil := make(chan string, 2)

	var wg sync.WaitGroup

	wg.Add(1)
	go worker(jobs, hasil, &wg)

	jobs <- 5
	jobs <- 10
	close(jobs)

	go func() {
		wg.Wait()
		close(hasil)
	}()

	for res := range hasil {
		fmt.Println(res)
	}

	fmt.Println("selesai...")

}
