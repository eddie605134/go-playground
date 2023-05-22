package varTest

import (
	"fmt"
	"strconv"
)

// 在變數名稱後加上變數型別，若無賦值(初始化)則變數會被初始化成其型別的預設值。
// 將型別放在變數名稱後方

/*
strconv :
strconv 是 Go 語言的標準函式庫之一，提供了字串和基本資料類型之間的相互轉換功能。它包含了許多函式，用於將字串解析為基本資料類型，或將基本資料類型轉換為字串。

以下是 strconv 包提供的一些常用函式和功能：

字串轉換為基本資料類型：

Atoi(s string) (int, error)：將字串轉換為 int 型態。
ParseFloat(s string, bitSize int) (float64, error)：將字串轉換為 float64 型態。
ParseInt(s string, base int, bitSize int) (int64, error)：將字串轉換為指定進制的整數。
ParseUint(s string, base int, bitSize int) (uint64, error)：將字串轉換為無符號整數。
ParseBool(str string) (bool, error)：將字串轉換為 bool 型態。
基本資料類型轉換為字串：

Itoa(i int) string：將 int 型態轉換為字串。
FormatFloat(f float64, fmt byte, prec, bitSize int) string：將 float64 型態轉換為字串。
FormatInt(i int64, base int) string：將整數轉換為指定進制的字串。
FormatUint(i uint64, base int) string：將無符號整數轉換為指定進制的字串。
AppendBool(dst []byte, b bool) []byte：將 bool 型態追加到位元組切片中。

*/
/*
fmt 是 Go 语言的一个标准库，提供了格式化输入和输出的功能。它包含了许多函数和方法，用于在控制台、文件和其他输出流中进行格式化输出。

以下是 fmt 包提供的一些常用功能：

格式化输出：

Print 和 Println：将给定的参数打印到标准输出，并自动添加空格和换行符。
Printf：使用格式化字符串，将参数按照指定的格式输出到标准输出。
Sprint 和 Sprintln：将给定的参数格式化为字符串，而不是直接输出到标准输出。
格式化输入：

Scan 和 Scanln：从标准输入读取数据，并将其存储到给定的参数中。
Scanf：从标准输入按照指定的格式读取数据，并将其存储到给定的参数中。
格式化字符串：

Sprintf：使用格式化字符串，将参数按照指定的格式转换为字符串。
Fprintf：使用格式化字符串，将参数按照指定的格式输出到指定的文件或输出流中。
其他功能：

Errorf：使用格式化字符串，创建一个新的错误对象。
Fprint 和 Fprintln：将给定的参数按照指定的格式输出到指定的文件或输出流中。
Scanf 和 Fscanf：从指定的文件或输入流按照指定的格式读取数据，并将其存储到给定的参数中。
fmt 包提供了一种简单而灵活的方式来格式化输入和输出，使得程序的输入输出更具可读性和可维护性。

*/

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
