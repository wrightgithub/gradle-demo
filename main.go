package main

import (
	_ "fmt"
	_ "github.com/urfave/cli"
	_ "os/exec"
	"os"
	"github.com/urfave/cli"
	log "github.com/Sirupsen/logrus"
)


func main() {

	//fmt.Println("helloxx")
	////log.Error("fuck")
	//c := make(chan int) // 分配一个goroutine
	//go func() {
	//	fmt.Println("test go")
	//	c <- 1
	//}()
	//res := <-c
	//
	//defer func() {fmt.Println("defer  will run finally....")}()
	//fmt.Printf("res=%d \n", res)
	//fmt.Println("---------------------------")
	//cmd := exec.Command("ls")
	//cmd.Stdout = os.Stdout
	//cmd.Run();
	//cmd.Wait();

	app := cli.NewApp()
	app.Name = "demo"
	app.Usage = "test command"

	app.Commands = []cli.Command{
		echoCommand,
	}


	app.Before = func(context *cli.Context) error {
		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})

		log.SetOutput(os.Stdout)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
