package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("== Паттерн Worker Pool  ==")

	numWorkers := 3
	numJobs := 15

	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workerPool(i, jobs, results, &wg)
	}

	go func() {
		defer close(jobs)
		for i := 1; i <= numJobs; i++ {
			fmt.Printf("Отправка джобы %d в пул\n", i)
			jobs <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Printf("\nПул обработчиков с %d обработчиками обрабатывает %d джобов: \n", numWorkers, numJobs)
	fmt.Println()

	count := 0
	for result := range results {
		fmt.Printf("Результат: %s\n", result)
		count++
	}

	fmt.Printf("\nПулл обработчиков завершен! Обработано %d джобов. \n", count)
}

func workerPool(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Обработчик %d стартовал\n", id)

	for job := range jobs {
		processingTime := time.Duration(rand.Intn(300)+200) * time.Millisecond
		fmt.Printf("Обработчик %d обрабатывает работу %d (заёмёт %v)\n", id, job, processingTime)

		time.Sleep(processingTime)

		result := fmt.Sprintf("Работа завершения обработчиком %d в %v", job, id, processingTime)
		results <- result
	}

	fmt.Printf("Обработчик %d завершился\n", id)
}
