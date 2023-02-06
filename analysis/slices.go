package analysis

import "fmt"

func Slices() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 5)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = slice[1:]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
