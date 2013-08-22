package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"mygos/mytesting"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	fmt.Println("begin :")

	mytesting.PrintBeegoVersion()
	mytesting.MainTest()
	fmt.Println("----------")
	beego.Router("/", &MainController{})
	beego.Run()
}
