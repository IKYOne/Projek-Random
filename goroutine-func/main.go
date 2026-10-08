package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	users := []string{"andi", "eko", "siti"}
	var wg sync.WaitGroup

	fmt.Println("Mulai megirim email...")

	for _, user := range users {
		wg.Add(1)

		go func(usr string) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Email berhasil terkirim ke: %s\n", usr)
		}(user)
	}

	for _, user := range users {
		wg.Add(1)

		usr := user
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Email berhasil terkirim ke: %s\n", usr)
		}()
	}

	wg.Wait()

}
