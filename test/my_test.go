package test

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	res := SayHello()
	fmt.Println(res)
	if res != "hello" {
		//t.Error("error olll")
		t.FailNow()
	}
	fmt.Println("success")
}
