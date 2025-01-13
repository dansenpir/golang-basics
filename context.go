package main

import (
	"Basics/utils"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// SimulateDBQuery simulates a database query that takes some time
func SimulateDBQuery(ctx context.Context) (string, error) {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Query Result"
	}()

	select {
	case result := <-ch:
		return result, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// HandleRequest is an HTTP handler that uses context
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	result, err := SimulateDBQuery(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "Request timed out", http.StatusRequestTimeout)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	_, err = fmt.Fprintf(w, "Result: %s", result)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

// WithValue demonstrates using context to pass request-scoped values
func WithValue(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "John Doe")
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// ValueAwareHandler is a handler that uses the value from context
func ValueAwareHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(string)
	if !ok {
		user = "anonymous user"
	}
	_, err := fmt.Fprintf(w, "Hello, %s!", user)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func demonstrateContextTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	result, err := SimulateDBQuery(ctx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("Query timed out")
		} else {
			fmt.Printf("Query error: %v\n", err)
		}
	} else {
		fmt.Printf("Query result: %s\n", result)
	}
}

func demonstrateContextValue() {
	ctx := context.WithValue(context.Background(), "user", "John Doe")
	user, ok := ctx.Value("user").(string)
	if !ok {
		user = "anonymous user"
	}
	fmt.Printf("Hello, %s!\n", user)
}

func demonstrateManualCancellation() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine: Cancelled")
			return
		case <-time.After(5 * time.Second):
			fmt.Println("Goroutine: Completed")
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second) // Give some time for the goroutine to print
}

func main() {
	var wg sync.WaitGroup

	utils.PrintSection("1. Context with Timeout")
	wg.Add(1)
	go func() {
		defer wg.Done()
		demonstrateContextTimeout()
	}()

	// Wait for the timeout demonstration to complete
	wg.Wait()

	utils.PrintSection("2. Context with Value")
	demonstrateContextValue()

	utils.PrintSection("3. Manual Context Manipulation")
	wg.Add(1)
	go func() {
		defer wg.Done()
		demonstrateManualCancellation()
	}()

	// Wait for the manual cancellation demonstration to complete
	wg.Wait()

	// Set up HTTP handlers
	http.HandleFunc("/query", HandleRequest)
	http.HandleFunc("/greet", WithValue(ValueAwareHandler))

	// Start the HTTP server in a separate goroutine
	go func() {
		fmt.Println("Starting server on :8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Give some time for the server to start
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\nServer is running. Press Ctrl+C to stop.")

	// Keep the main function running
	select {}
}
