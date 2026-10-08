package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

type Result struct {
	jobID  int
	Output string
	Err    error
}


type Res

func worker(id int, jobs <-chan int, hasil chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {

		fmt.Printf("👷 Worker %d: Memproses Job %d...\n", id, job)
		time.Sleep(500 * time.Millisecond)

		if job == 3 {
			hasil <- Result{
				jobID: job,
				Err:   fmt.Errorf("gagal memproses data %d: koneksi database terputus", job),
			}
			continue
		}

		hasil <- Result{
			jobID:  job,
			Output: fmt.Sprintf("Data Job %d berhasil diolah", job),
			Err:    nil,
		}
	}
}

func GetDatabase(ctx context.Context, db *sql.DB, wg *sync.WaitGroup) () {
	defer wg.Done()

	query := "SELECT * FROM mahasiswa"

	data, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil
	}
}

func main() {
	totatJobs := 10
	totalWorker := 5
	job := make(chan int, totatJobs)
	resolut := make(chan Result, totalWorker)
	var wg sync.WaitGroup

	for w := range totalWorker {
		wg.Add(1)
		go worker(w, job, resolut, &wg)
	}

	for j := range totatJobs {
		job <- j
	}

	close(job)
	go func() {
		wg.Wait()
		close(resolut)
	}()

	for res := range resolut {

		if res.Err != nil {
			fmt.Printf("[Error] job %d: %v\n", res.jobID, res.Err)
		} else {
			fmt.Printf("[Success] Job %d: %s\n", res.jobID, res.Output)
		}
	}

}
