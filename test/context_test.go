package test

import (
	"context"
	"testing"
	"time"
	"fmt"
)

type User struct {
	UserID string
	UserName string
}

type Context interface {
	Done() <- chan struct{}
	Err() error
}

func TestContextWithValue (t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var userData = User{UserID: "AAA", UserName:"AAA_NAME"}
	ctx = context.WithValue(ctx, "current_user",userData)
	time.Sleep(15*time.Second)

	var currentUser User

	if v := ctx.Value("current_user"); v != nil {
		u, ok := v.(User)
		if !ok {
			fmt.Println("current user : not available!")
		} 
		currentUser = u
		fmt.Println("current user : ", currentUser)
	} else {
		fmt.Println("current user : not available!")
	}
}

func TestContextWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func(){
		cancel()
	}()

	result, err := workProcessWithContext(ctx)

	if err != nil {
		fmt.Println("error :", err)
	}

	fmt.Println("result :", result)
}

func workProcessWithContext(ctx context.Context) (string,error) {
	done := make(chan string)

	go func() {
		done <- workProcess()
	}()

	select {
	case result := <- done:
		return result, nil
	case <-ctx.Done():
		return "Failed from context message", ctx.Err()
	}
}

func workProcess() string {
	<-time.After(3*time.Second)
	return "Done"
}