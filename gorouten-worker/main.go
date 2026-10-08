package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, resoult chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()


	for job := range jobs {
		fmt.Printf("👷 [Worker %d] Mulai memproses Job #%d\n", id, job)


		time.Sleep(100 * time.Millisecond)
		resoult <- fmt.Sprintf("Job #%d SUKSES dikerjakan oleh Worker %d", job, id)
	}

	fmt.Printf("🛑 [Worker %d] Berhenti (Antrean kosong).\n", id)
}



func worker1()

func main() {

	totalJob := 10
	totalWorkers := 3

	jobs := make(chan int, totalJob)
	results := make(chan string, totalJob)

	var wg sync.WaitGroup

	for w := 1; w <= totalWorkers; w++ {
		wg.Add(1)

		go
	}
}
