package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "cjw",
		"course": "golang",
	}
	m2 := make(map[string]string)
	var m3 map[string]string
	fmt.Println(m, m2, m3)
	for k, v := range m {
		fmt.Println(k, v)
	}
	//获取数据
	if coursename, ok := m["course"]; ok {
		fmt.Println(coursename)
	} else {
		fmt.Println("key 'cause' does not exist")
	}
	//删除数据
	delete(m, "name")
	s, ok := m["name"]
	fmt.Println(s, ok)
}
