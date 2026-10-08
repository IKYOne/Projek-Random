package main

import (
	"fmt"
	"sync"
)

func makanan(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("Pesan makanan #%d", i)
	}
}

func minuman(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("Pesan Minuman #%d", i)
	}
}

func listener(makanan <-chan string, minuman <-chan string, quit <-chan any) {
	fmt.Println("👨‍🍳 [DAPUR] Dapur siap menerima pesanan...")
	fmt.Println("--------------------------------------------------")
	for {
		select {
		case msgMakanan := <-makanan:
			fmt.Println("[MAKANAN]", msgMakanan)
		case msgMinuman := <-minuman:
			fmt.Println("[MINUMAN]", msgMinuman)
		case <-quit:
			fmt.Println("--------------------------------------------------")
			fmt.Println("🔔 [DAPUR] Semua pesanan selesai dibuat! Restoran tutup.")
			return
		}
	}
}

func main() {

	var wg sync.WaitGroup

	makananCh := make(chan string)
	minumanCh := make(chan string)
	quitCh := make(chan any)
	wg.Add(2)
	go makanan(makananCh, &wg)
	go minuman(minumanCh, &wg)

	go listener(makananCh, minumanCh, quitCh)

	go func() {
		wg.Wait()
		close(quitCh)
	}()

	<-quitCh
	fmt.Println("Main Selesai.")

}
