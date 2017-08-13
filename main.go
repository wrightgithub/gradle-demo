package main

import (
	"fmt"
	_ "github.com/urfave/cli"
	"os/exec"
	"os"
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
	fmt.Println("---------------------------")
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	cmd.Run();
	cmd.Wait();
}
