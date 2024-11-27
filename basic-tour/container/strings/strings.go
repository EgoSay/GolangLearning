package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello,陈嘉伟"
	for _, b := range []byte(s) {
		//十六进制输出字节
		fmt.Printf("%X", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		//函数解码开始位置的第一个utf-8编码的码值，返回该码值和编码的字节数
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
}
