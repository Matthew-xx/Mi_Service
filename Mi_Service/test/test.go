package main

import (
	"fmt"
	pd "pro_package"
)
func main()  {
	text := pd.Person{
		Name:"maxx",
		Age:21,
		Email:"1659530762@qq.com",
		Reps:[]int64{18,24,27},
	}
	fmt.Println(text)
}
