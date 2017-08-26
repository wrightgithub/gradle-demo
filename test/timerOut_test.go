package test

import (
	"time"
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	//closeChannel()
	c := make(chan int)

	// 都只执行一次
	timeout := time.After(time.Second * 2) //
	t1 := time.NewTimer(time.Second * 3)   //
	var i int
	go func() {
		for {
			select {
			case x, ok := <-c:
				fmt.Println("channel sign,", x, ok)
				return
			case <-t1.C: // 代码段2
				fmt.Println("3s定时任务")
			case <-timeout: // 代码段1
				i++
				fmt.Println(i, "2s定时输出")
			case <-time.After(time.Second * 4): // 代码段3  有default存在  这个超时后回去执行default
				fmt.Println("4s timeout。。。。")
			default: // 代码段4
				fmt.Println("default")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	select {}
	//time.Sleep(time.Second * 6)
	//close(c)
	time.Sleep(time.Second * 2)
	fmt.Println("main退出")
}

//发送者
func sender(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
		if i >= 5 {
			time.Sleep(time.Second * 7)
		} else {
			time.Sleep(time.Second)
		}
	}
}

func TestSelectTimer(t *testing.T) {
	c := make(chan int)
	// 心跳
	heartbeat := time.Tick(5 * time.Second)
	go sender(c)
	timeout := time.After(time.Second * 3)
	for {
		select {
		case d := <-c:
			fmt.Println(d)
		case <-timeout:
			fmt.Println("这是定时操作任务 >>>>>")
		case dd := <-time.After(time.Second * 3):
			fmt.Println(dd, "这是超时*****")
		case <-heartbeat:
			fmt.Println("this is heartbeat ")
		}

		fmt.Println("for end")
	}
}
