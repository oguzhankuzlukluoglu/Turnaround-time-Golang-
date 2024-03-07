package main

import (
	"fmt"
	"time"
)

func main() {
	performComplexTimeTracking()
}

func performComplexTimeTracking() {

	startTime := time.Now()
	defer func() {

		endTime := time.Now()
		duration := endTime.Sub(startTime)
		fmt.Printf("Fonksiyon çalışma süresi: %v\n", duration)
	}()

	for i := 0; i < 3; i++ {
		fmt.Printf("İşlem %d başladı.\n", i+1)

		time.Sleep(time.Second)

		fmt.Printf("İşlem %d tamamlandı.\n", i+1)
	}

	performAdditionalTask()

	fmt.Println("Fonksiyon çalıştı.")
}

func performAdditionalTask() {

	startTime := time.Now()
	defer func() {

		endTime := time.Now()
		duration := endTime.Sub(startTime)
		fmt.Printf("İç fonksiyon çalışma süresi: %v\n", duration)
	}()

	for i := 0; i < 2; i++ {
		fmt.Printf("İç işlem %d başladı.\n", i+1)

		time.Sleep(500 * time.Millisecond)

		fmt.Printf("İç işlem %d tamamlandı.\n", i+1)
	}
}
