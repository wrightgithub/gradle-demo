package test

import (
	"fmt"
	"testing"
	"time"
)

// 定义一个结构sharded_var作为一个共享变量
//共享变量有一个读通道和一个写通道组成
type sharded_var struct {
	reader chan int
	writer chan int
}

//共享变量维护协程
// 0 chan通道其实就是读一个放一个，没有任何buffer，所有写都阻塞
func sharded_var_whachdog(v sharded_var) {

	go func() {

		//初始值
		var value int = 0

		for {

			//监听读写通道，完成服务

			select {

			case value = <-v.writer:
				// 因为默认大小为0 ，所以一开始就阻塞在这里，
			case v.reader <- value:
				fmt.Println("read")

			}

		}

	}()

}

func TestShare(t *testing.T) {

	//初始化，并开始维护协程

	v := sharded_var{make(chan int, 0), make(chan int, 0)}

	sharded_var_whachdog(v)

	//读取初始值

	//fmt.Println(<-v.reader)

	//写入一个值
	v.writer <- 1

	//读取新写入的值
	fmt.Println(<-v.reader)
	time.Sleep(time.Second * 10)

}
