package main

import "fmt"

func main() {
	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("%v len:%d ptr:%p\n", a, len(a), cap(a), a)
	}

}
