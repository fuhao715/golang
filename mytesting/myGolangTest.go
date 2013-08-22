package mytesting

import (
	"fmt"
	"math"
	"mygos/myError"
	"mygos/myGolang"
)

// var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面
var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no!"

const (
	Pi = 3.14
)

func add(x int, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	fmt.Println("x = ", x)
	y = sum - x
	fmt.Println("y = ", y)
	return
}

func MainTest() {
	fmt.Println("Hello world，世界 ! \n ", math.Pi, "Day")
	fmt.Printf("Now you have %g problems .", math.Nextafter(2, 3))
	// 新增函数教程
	fmt.Println("\n ", add(42, 13))
	// 函数多返回值
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	// 如果命名了返回值参数，一个没有参数的 return 语句，会将当前的值作为返回值返回。
	fmt.Println(split(17))
	//  设置变量
	fmt.Println(x, y, z, c, python, java)
	var r, q, t int = 4, 5, 6
	golang, ruby, perl := true, false, "no!"
	fmt.Println(r, q, t, golang, ruby, perl)

	const World = "世界"
	fmt.Println("hello ", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println("end the world ,再见第一课golang test 乱码状态 github 是否有乱码呢？版本状态")
	fmt.Println("ideavim test 非常好用\n")
	fmt.Println("myGolang is ", myGolang.HelloWorld("fuhao"))
	fmt.Println(myGolang.Sqrt(4.0))
	fmt.Println(myGolang.Compute_circle(1.4))
	fmt.Println(myGolang.PI)
	myGolang.TestBoolean()
	myGolang.Complex()
	myGolang.Splice("splice some words ! ")
	myGolang.Array_Test()
	myGolang.MyPrint()
	val, err := myGolang.Multi_ret("one")
	fmt.Printf("val = %d \t err = %v \n", val, err)
	myGolang.MyPointer()
	myGolang.MyMap()
	myGolang.New_Make()
	myGolang.Test_Multi_argus()
	myGolang.Test_Closure()
	x1, y1 := myGolang.Split(24)
	fmt.Printf("x = %d  ,y = %d \n ", x1, y1)
	fmt.Println(myGolang.NeedInt(myGolang.Small))
	fmt.Println(myGolang.NeedFloat(myGolang.Small))
	fmt.Println(myGolang.NeedFloat(myGolang.Big))
	// fmt.Println(myGolang.NeedInt(myGolang.Big))
	myError.LogError()

	fmt.Println("begin :")
	fmt.Println("----------")
}
