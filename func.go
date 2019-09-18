package go_study

import "fmt"

type cb func(num int) int

type ChipType int

const (
	None ChipType = iota
	CPU           // 中央处理器
	GPU           // 图形处理器
)

//方法 相当于类ChipType的成员方法
func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func main() {
	// 输出CPU的值并以整型格式显示
	fmt.Printf(CPU.String())
	fmt.Printf("%s %d", CPU, CPU)
	x := 12
	y := 13
	x, y = swap(x, y)
	fmt.Println(x)
	fmt.Println(y)

	num := testCb(12, callback)
	fmt.Println(num)

	f := closure(1)
	fmt.Println(f(2))          //3
	fmt.Println(testCb(12, f)) //13
}

func callback(num int) int {
	num++
	return num
}

func testCb(num int, callback cb) int {
	return callback(num)
}

func closure(num int) func(int) int {
	return func(num1 int) int {
		return num1 + num
	}
}

func swap(x, y int) (int, int) {
	return y, x
}

func init() {
	fmt.Printf("init")
}
