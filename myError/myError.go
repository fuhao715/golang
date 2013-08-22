/**
 * Created with IntelliJ IDEA.
 * User: fuhao
 * Date: 13-7-13
 * Time: 下午5:17
 * To change this template use File | Settings | File Templates.
 */
package myError

import (
	"errors"
	"fmt"
)

func LogError() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}
