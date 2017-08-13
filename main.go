package main

import (
	"fmt"
	_ "github.com/urfave/cli"
)

func main() {

	fmt.Println("helloxx")
	//log.Error("fuck")
	c := make(chan int) // 分配一个goroutine
	go func() {
		fmt.Println("test go")
		c <- 1
	}()
	res := <-c

	fmt.Printf("res=%d \n", res)
}
