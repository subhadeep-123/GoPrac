package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "api-key", "test-api-key")
}
func doSomethingCool(ctx context.Context) {
	apiKey := ctx.Value("api-key")
	if apiKey != nil {
		fmt.Println(apiKey)
	}
}

func doSometingEvenCooler(ctx context.Context) {
	// conext with deadlines
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			err := ctx.Err()
			fmt.Println(err)
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
func main() {
	fmt.Println("go Context tutorial")
	ctx := context.Background()
	ctx = enrichContext(ctx)
	doSomethingCool(ctx)

	fmt.Printf("==============================================\n\n")

	// context with deadlines
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go doSometingEvenCooler(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("os no, I've exceeded the deadline")
	}
	time.Sleep(2 * time.Second)
}
