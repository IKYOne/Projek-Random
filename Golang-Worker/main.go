package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, resoult chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("data akan di kerjakan...")
		time.Sleep(300 * time.Millisecond)
		resoult <- fmt.Sprintf("data %d dan tugas: %d\n", id, job)
		fmt.Println("data selesai di kerjakan")
	}

	fmt.Println("data habis dikerjakan & dan keluar")
}

func main() {
	totalJob := 10
	totalWorker := 10

	job := make(chan int, totalJob)
	resoult := make(chan string, totalWorker)

	var wg sync.WaitGroup

	for w := 1; w <= totalWorker; w++ {
		wg.Add(1)
		go worker(1, job, resoult, &wg)
	}

	for j := 1; j <= totalJob; j++ {
		job <- j
	}

	close(job)

	go func() {
		wg.Wait()
		close(resoult)
	}()

	for res := range resoult {
		fmt.Printf("data %s", res)
	}

}
