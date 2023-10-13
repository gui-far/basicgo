package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/*
func soma(x int, y int) (int, bool) {
	if x > 10 {
		return x + y, true
	}

	return x + y, false
}
*/

// Like JS Obeject or Class
// First Capitalized means Public, otherwise is Privated
type Course struct {
	Name        string `json:"course"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {

	//Example 1
	/*
		course := Course{
			Name:        "Golang",
			Description: "Golang Course",
			Price:       100,
		}

		course.Name = "Golang2"

		fmt.Println(course.GetFullInfo())

		http.HandleFunc("/", home)
		http.ListenAndServe(":8080", nil)
	*/

	//Example 2 - Threads
	//go counter()
	//go counter()
	//counter()

	//Example 3 - Channels (channels avoid race conditions)
	/*
		channel := make(chan string)

		go func() {
			channel <- "Hello"
		}()

		fmt.Println(<-channel)
	*/

	//Example 4 - Worker
	channel := make(chan int)
	//go worker(1, channel)
	//go worker(2, channel)
	//go worker(3, channel)

	for i := 1; i <= 100000; i++ {
		go worker(i, channel) //2kb
	}

	for i := 0; i < 100000; i++ {
		channel <- i
	}
}

func home(w http.ResponseWriter, r *http.Request) {

	course := Course{
		Name:        "Golang",
		Description: "Golang Course",
		Price:       100,
	}

	//res, _ := json.Marshal(course)
	json.NewEncoder(w).Encode(course)

	w.Write([]byte("Hello World"))
}
