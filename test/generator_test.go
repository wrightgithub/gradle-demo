package test

import (
	"fmt"
	"math/rand"
	"testing"
	_ "time"
	"time"
)

// 函数rand_generator_2，返回通道(Channel)

func rand_generator_2() chan int {

	// 创建通道

	out := make(chan int, 10)

	// 创建协程

	go func() {

		for {

			time.Sleep(time.Millisecond * 5)
			//向通道内写入数据，如果无人读取会等待
			out <- rand.Int()

		}

	}()

	return out
}

func TestGenerator(t *testing.T) {

	// 生成随机数作为一个服务

	rand_service_handler := rand_generator_2()

	// 从服务中读取随机数并打印

	for {
		select {
		case ret := <-rand_service_handler:
			fmt.Printf("%d\n", ret)
			return
		default:
			// do other thing
		}
	}

}

func rand_generator_3(num int) chan int {

	//创建通道

	ret := make(chan int, 10)

	for i := 0; i < num; i++ {
		go func() {
			for {
				ret <- <-rand_generator_2()
			}
		}()
	}
	return ret

}

func TestGenerator3(t *testing.T) {

	//千万不要这样写
	//	for i := 0; i < 100; i++ {
	//	<-rand_generator_3()
	//}
	//因为这样写每次都会 执行一遍rand_generator_3里面的for，每次都执行一遍函数，肯定会慢
	//只要拿到chan的引用就可以了

	ret := rand_generator_3(10);
	t1 := time.Now()
	for i := 0; i < 100; i++ {
		<-ret
	}

	fmt.Println("App elapsed: ", time.Since(t1))
}
