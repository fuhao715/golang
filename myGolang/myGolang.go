/**
 * Created with IntelliJ IDEA.
 * User: lenovo
 * Date: 13-7-8
 * Time: 下午10:04
 * To change this template use File | Settings | File Templates.
 */
package myGolang
import (
	"fmt"
	"math"
)
var isActive bool         // 全局变量申明
var enabled , disabled = true ,false    // 忽略类型的申明

func TestBoolean() {
	isActive = false
	fmt.Println(isActive)
	var isAvailable bool //  一般申明
	isAvailable = true   // 赋值操作
	valid := false      // 简短申明
	fmt.Println(isAvailable)
	fmt.Println(enabled,disabled,valid)
}

func HelloWorld(who string) string{
	fmt.Printf("Hello world，世界 %s ",who)
	return who
}

func Sqrt(x float64) float64 {
	z := 0.0
	fmt.Println("input args is ",x)
	for i :=0 ; i <100 ; i++ {
		z -= (z*z - x ) / (2 * x)
	}
	return z
}
const (
	PI = 3.1415926
	con = 77
)

const(
	x = iota  // x == 0
	y = iota  // y == 1
	z = iota  // z == 2
	w  // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	e, f, g = iota, iota, iota //e=0,f=0,g=0 iota在同一行值相同
)

func Compute_circle(r float64) float64 {
	return PI*r*r
}

func Complex() {
	var c complex64 = 5 + 4i
	fmt.Printf("Value is %v \n",c)
}

func Splice(s string) {
	var sp = `hello ` +" splice the  " + s[1:]
	var moreLineStr = `firest line
    seconde line
	third line
	`
	fmt.Printf("sp is %v \n %s \n",sp,moreLineStr)
	fmt.Printf("x= %d,y = %d ,z = %d, w=  %d,v = %d ,e =  %d ,f=%d ,g = %d \n ",x,y,z,w,v,e,f,g)
}

// Array

func Array_Test() {
	var arr [10]int  // 声明了一个int类型的数组
	arr[0] = 42      // 数组下标是从0开始的
	arr[1] = 13      // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0])  // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9]) //返回未赋值的最后一个元素，默认返回0

	a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组

	b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0

	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

	fmt.Printf("a %d ,b %d , c %d \n ",a,b,c)
	fmt.Printf("a len %d ,b len  %d , c len  %d \n ",len(a),len(b),len(c))

	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Printf("doubleArray %d ,easyArray %d  ",doubleArray,easyArray)
}

func MyPrint(){
	fmt.Println("hello world 中国golang社区")
	fmt.Printf("%t \n",1==2)
	fmt.Printf("二进制：%b \n",255)
	fmt.Printf("八进制:%o \n",255)
	fmt.Printf("十六进制:%x \n",255)
	fmt.Printf("十进制:%d \n",255)
	fmt.Printf("浮点数:%f \n",math.Pi)
	fmt.Printf("字符串:%s \n","hello world")
}

func MyMap(){
	m := make(map[string]int) //使用make创建一个空的map

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m) //输出 map[three:3 two:2 one:1] (顺序在运行时可能不一样)
	fmt.Println(len(m)) //输出 3

	v := m["two"] //从map里取值
	fmt.Println(v) // 输出 2

	delete(m, "two")
	fmt.Println(m) //输出 map[three:3 one:1]

	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(m1) //输出 map[two:2 three:3 one:1] (顺序在运行时可能不一样)

	for key, val := range m1{
		fmt.Printf("%s => %d \n", key, val)
		/*输出：(顺序在运行时可能不一样)
			three => 3
			one => 1
			two => 2*/
	}
}

func MyPointer(){
	var i int = 1
	var pInt *int = &i
	//输出：i=1     pInt=0xf8400371b0       *pInt=1
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)

	*pInt = 2
	//输出：i=2     pInt=0xf8400371b0       *pInt=2
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)

	i = 3
	//输出：i=3     pInt=0xf8400371b0       *pInt=3
	fmt.Printf("i=%d\tpInt=%p\t*pInt=%d\n", i, pInt, *pInt)
}

func New_Make(){
	var p *[]int = new([]int)
	fmt.Println(p) //输出：&[]
	*p = make([]int, 10, 10)
	fmt.Println(p) //输出：&[0 0 0 0 0 0 0 0 0 0]
	fmt.Println((*p)[2]) //输出： 0

	// 习惯用法:
	v := make([]int, 10)
	fmt.Println(v) //输出：[0 0 0 0 0 0 0 0 0 0]
}

func Multi_ret(key string) (int,bool){
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	var err bool
	var val int

	val, err = m[key]

	return val, err
}

func Multi_argus(nums ...int){
	fmt.Println(nums)
	total :=0
	for _,num := range nums {
		total +=num
	}
	fmt.Println(total)
}

func Test_Multi_argus(){
	Multi_argus(1, 2)
	Multi_argus(1, 2, 3)

	//传数组
	nums := []int{1, 2, 3, 4}
	Multi_argus(nums...)
}

func Closure() func() int{
	i,j := 1,1
	return func() int {
		var tmp = i +j
		i,j = j,tmp
		return tmp
	}
}

func Test_Closure(){
	nextNumFunc := Closure()
	for i:=0;i<10;i++{
		fmt.Println(nextNumFunc())
	}
}

func Split(sum int)(x,y int ){
	x = sum * 4 / 9
	y = sum -x
	return
}

const (
	Big = 1 << 100
	Small = Big >> 99
)
func NeedInt(x int ) int {
	return x * 10 +1
}

func NeedFloat(x float64) float64{
	return x * 0.1
}
