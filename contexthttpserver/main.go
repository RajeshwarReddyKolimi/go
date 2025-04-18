package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func longRunningTask(ch chan bool) {
	fmt.Println("started")
	time.Sleep(5 * time.Second)
	fmt.Println("ended")
	ch <- true
}
func handleProcess(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	ch := make(chan bool)

	go longRunningTask(ch)

	select {
	case <-ctx.Done():
		err := ctx.Err()
		switch err {
		case context.DeadlineExceeded:
			fmt.Println("Request cancelled due to timeout")
			http.Error(w, "Request timeout", http.StatusGatewayTimeout)
		case context.Canceled:
			fmt.Println("Request cancelled by client")
			http.Error(w, "Request cancelled", http.StatusRequestTimeout)
		}
	case <-ch:
		fmt.Println("Task completed")
	}
}
func main() {
	http.HandleFunc("/process", handleProcess)
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
