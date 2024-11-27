/**
 * @Author codeAC
 * @Time: 2018/8/9 11:00
 * @Description  defer相当于Java中的finally
 */

package main

import (
	"Learning/functional/Fibonacci"
	"bufio"
	"fmt"
	"os"
)

func WriteFile(filename string) {
	file, err := os.OpenFile(filename,
		os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(pathError)
		} else {
			fmt.Printf("%s,%s,%s",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
	}
	//关闭文件
	defer file.Close()
	writer := bufio.NewWriter(file)
	//将字符流刷新到文件中
	defer writer.Flush()
	fib := Fibonacci.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, fib())
	}
}
func main() {
	WriteFile("test.txt")
}
