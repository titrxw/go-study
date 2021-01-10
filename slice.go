package go_study

import "fmt"

//https://blog.csdn.net/ahilll/article/details/84624120

func testAppend(test []int) {
	test[0] = 3
	test = append(test, 2)
	test[0] = 4
	fmt.Println(test) //[4,2]
}

func newSlice() {
	tmp := []int{1, 2, 3, 4}
	fmt.Println(len(tmp))
	fmt.Println(cap(tmp))
	//4 4

	tmp1 := make([]int, 3, 5)
	fmt.Println(len(tmp1))
	fmt.Println(cap(tmp1))
	//3 5

	tmp1[0] = 1
	tmp1[1] = 1
	tmp1[2] = 1
	fmt.Println(len(tmp1))
	fmt.Println(cap(tmp1))
	//3 5

	tmp2 := append(tmp1, 1)
	fmt.Print(tmp1)
	fmt.Print(tmp2)
	fmt.Println(len(tmp2))
	fmt.Println(cap(tmp2))
	//如果len(a)+1<=cap(a)，这个时候a内部的数组仍足够存储新添加的数据，此时，b的ptr和a的ptr是相同的。此时：b.ptr == a.ptr, len(b) == len(a)+1, cap(b) == cap(a)。
	//append()后，s1和s2确实指向了不同的底层数组，并且二者的数组容量cap也不相同了。过程是这样的:当append()时，发现数组容量不够用，于是开辟了新的数组空间，cap变为原来的两倍，s2指向了这个新的数组，所以当修改s2时，s1不受影响
	//4 5

	tmp1 = append(tmp1, 1)
	fmt.Println(len(tmp1))
	fmt.Println(cap(tmp1))
	//5 5
	tmp1 = append(tmp1, 1)
	fmt.Println(len(tmp1))
	fmt.Println(cap(tmp1))
	//6 10
	//https://www.cnblogs.com/f-ck-need-u/p/9854932.html
}

func main() {
	newSlice()

	test12 := []int{1}
	fmt.Println(test12) //[1]
	testAppend(test12)
	fmt.Println(test12) //[3]
}
