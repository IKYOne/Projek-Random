package main

import (
	"fmt"
	"sync"
)

type MahasiswaPesan struct {
	Status  string
	Message string
}

func kirim(ch chan<- MahasiswaPesan, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)

	ch <- MahasiswaPesan{
		Status:  "Success",
		Message: "Data dikirim dari kelas A",
	}

	ch <- MahasiswaPesan{
		Status:  "Success",
		Message: "Data dikirim dari kelas B",
	}

}

func terima(ch <-chan MahasiswaPesan, wg *sync.WaitGroup) {
	defer wg.Done()

	for pesan := range ch {
		fmt.Printf("pesan: %s\n", pesan)
	}

}

func main() {
	var wg sync.WaitGroup

	ch := make(chan MahasiswaPesan, 2)

	wg.Add(2)
	go kirim(ch, &wg)
	go terima(ch, &wg)

	wg.Wait()

	fmt.Println("goroutine selesai...")

}
