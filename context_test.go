package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Context Background
func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

// Context with Value
func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// Context Get Value
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextA.Value("b"))

}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // slow simulation
			}
		}
	}()
	return destination
}

// Context with Cancel

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Goroutine Total: ", runtime.NumGoroutine())

	parent := context.Background()
	// ctx, cancel := context.WithCancel(parent) //Context with cancel
	ctx, cancel := context.WithTimeout(parent, 5*time.Second) // Context with timeout
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Goroutine total: ", runtime.NumGoroutine())
}

// Context with Deadline

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Goroutine Total: ", runtime.NumGoroutine())

	parent := context.Background()
	// ctx, cancel := context.WithCancel(parent) //Context with cancel
	// ctx, cancel := context.WithTimeout(parent, 5*time.Second) // Context with timeout
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second)) // Context with Deadline
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Goroutine total: ", runtime.NumGoroutine())
}
