package test

import (
	"context"
	"fmt"
	"testing"
	"time"
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

func TestChannel(t *testing.T) {

	testChan := make(chan string)

	go Chantt(testChan)

	fmt.Println("chan value :", <-testChan)
}

func Chantt(channel chan string) {
	time.Sleep(10*time.Second)
	channel <- "OK"
}

func TestPanic(t *testing.T) {
	fmt.Println(divide(1,0))
}

func divide(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	return a / b
}

func TestRaceDetector(t *testing.T) {
	c := make(chan bool)
	m := make(map[string] string)

	go func(){
		m["12"] = "a"
		c <- true
	}()

	m["13"] = "b"
	<- c
	for k, v:= range m {
		fmt.Println(k,v)
	}

	//In terminal, you can text this "go run -race race.go"
}

func TestContextWithCancel3(t *testing.T) {
	ctx, ctx_cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(10*time.Second)
		ctx_cancel()
	}()

	select {
	case <- ctx.Done():
		fmt.Println("context canceled!")
	}
}

func TestContextWithTimeout(t *testing.T) {
	ctx, ctx_cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer ctx_cancel()

	go func() {
		time.Sleep(10*time.Second)
		fmt.Println("not working")
	}()

	select {
	case <- ctx.Done():
		fmt.Println("Timeout!! normal operating")
	}
}
