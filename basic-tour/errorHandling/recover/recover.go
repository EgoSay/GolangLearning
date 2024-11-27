/**
 * @Author codeAC
 * @Time: 2018/8/9 18:40
 * @Description  recover和panic
 */
package main

import (
	"errors"
	"fmt"
)

/**
 * 测试panic 和 recover
 * Go语言提供了recover内置函数，前面提到，一旦panic，逻辑就会走到defer那，那我们就在defer那等着，
 *调用recover函 数将会捕获到当前的panic（如果有的话），被捕获到的panic就不会向上传递了
 */
func tryRecover() {
	defer func() {
		r := recover()
		if r == nil {
			fmt.Println("Nothing is to Recover")
			return
		}
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()
	panic(errors.New("this is an error"))
	/*b := 0
	a := 5/b
	fmt.Println(a)*/
	//panic(123)
}
func main() {
	tryRecover()
}
