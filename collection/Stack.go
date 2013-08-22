/**
 * Created with IntelliJ IDEA.
 * User: lenovo
 * Date: 13-7-20
 * Time: 下午3:23
 * To change this template use File | Settings | File Templates.
 */
package collection

type Stack struct {
	data []interface {}
}

// 压栈顶
func (s *Stack) Push(x interface {}){
	s.data = append(s.data,x)
}

// 弹栈顶，并删除栈顶
// 如果是空栈将会引发运行时错误
func (s *Stack) Pop() interface {} {
	i := len(s.data) - 1
	res := s.data[i]
	s.data[i] = nil
	s.data = s.data[:i]
	return res
}

// 返回栈大小
func (s *Stack) Size() int {
	return len(s.data)
}


