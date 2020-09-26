package main


import (
	_ "HelloBeego0406/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("hello.go")
	beego.Run()
}
