package varTest

import (
	"fmt"
	"strconv"
)

// 在變數名稱後加上變數型別，若無賦值(初始化)則變數會被初始化成其型別的預設值。
// 將型別放在變數名稱後方

var age int
var enable bool

const Pi = 3.1415926
const Size = 50
const prefix = "Pluto_" // 常數的值在編譯時期就已經確定，所以可以用來定義陣列的長度

func VarTest() {
	name := "eddie"

	// var num = [4]int{2, 7, 11, 15}

	fmt.Printf("Hello %s\n", name)

	const (
		Taipei = iota
		Taoyuan
		Taichung
		Kaohsiung
		Tainan
	)

	fmt.Printf(strconv.Itoa(Taipei)) // 將整數轉換為字串
}
