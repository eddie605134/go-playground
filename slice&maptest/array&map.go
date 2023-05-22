package sliceAndMap

import "fmt"

func testSlice() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Print(slice[1:3]) // [2 3]
	fmt.Print(slice[1:])  // [2 3 4 5]
	fmt.Print(slice[:3])  // [1 2 3]
	fmt.Print(slice[:])   // [1 2 3 4 5]
}

func testMap() {

}
