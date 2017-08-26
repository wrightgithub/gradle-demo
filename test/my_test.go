package test

import (
	"fmt"
	"testing"
	"strings"
	"helloxyy.com/gradle-demo/utils"
)

func TestHello(t *testing.T) {
	res := utils.SayHello()
	fmt.Println(res)
	if res != "hello" {
		//t.Error("error olll")
		t.FailNow()
	}
	fmt.Println("success")
}


func TestDemo1(t *testing.T) {

	ret:=strings.Repeat("0", 1 << uint8(24 - 8))

	fmt.Println(string(ret[0]))
	fmt.Println('0')
}

