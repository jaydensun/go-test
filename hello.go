package main

import "fmt"

func main() {

	ss := make([]int, 5)
	ss = append(ss, 4)
	fmt.Println(ss)
	fmt.Println(&ss[0])
	println(ss)

}
