package main

import (
	"fmt"
	"net/http"
)

type student struct {
	Name string
	Age  int
}

func main() {

	_, err := http.Get("www.baidu.com")
	fmt.Println(err.Error())
	fmt.Println(err)

	m := make(map[string]*student)
	stus := []student{{"zs", 1},
		{"lisi", 2},
		{"wangwu", 3},
	}

	for index, s := range stus {
		fmt.Printf("%v,%v\n", index, s)
		m[s.Name] = &s
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("k:%v,v:%v\n", k, v)
	}

}
