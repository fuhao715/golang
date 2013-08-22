/**
 * Created with IntelliJ IDEA.
 * User: lenovo
 * Date: 13-7-20
 * Time: 下午3:33
 * To change this template use File | Settings | File Templates.
 */
package collection
import (
	collection "."
	"fmt"
)

func Example(){
	var s collection.Stack
	s.Push("world")
	s.Push("hello,")
	for s.Size() > 0 {
		fmt.Print(s.Pop())
	}
	fmt.Println()
}



