// main.go
package main

import (
	"fmt"
	"learninggolang/interfaces/infrastructure/postgres"
	"learninggolang/interfaces/service"
	"learninggolang/interfaces/utils"
	"sync"
)

type Order struct {
	ID    int
	Items []string
}

func handleOrder(order Order) {
	// process order
}

func main() {
	repo := postgres.NewPostgresUserRepo()
	userService := service.NewUserService(repo)

	userService.CreateUser("Sameer")
	utils.Contains([]int{1, 2, 3}, 2)
	utils.Contains([]string{"a", "b"}, "b")

	// simulate multiple API calls
	urls := []string{
		"https://api.example.com/data1",
		"https://api.example.com/data2",
		"https://api.example.com/data3",
	}
	// waitgroups
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			callAPI(u)
		}(url)
	}

	wg.Wait() // block until all done

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= 5; i++ {
		fmt.Println(<-results)
	}

}

// goroutines
func processOrders(orders []Order) {
	for _, order := range orders {
		order := order // avoid closure bug
		go func() {
			handleOrder(order)
		}()
	}
}

func callAPI(url string) {
	// simulate API call
}

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

//channels

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * 2
	}
}
