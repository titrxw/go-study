package main

import "fmt"

func main() {
	var te = [...]int{1, 2, 3, 4}
	fmt.Println(te[1])

	var arr [5]int
	arr = [5]int{1, 2, 3, 4, 5}
	arr[0] = 1
	arr[1] = 1
	arr[2] = 1
	arr[3] = 1
	arr[4] = 1

	var arr5 = arr
	arr5[1] = 6
	fmt.Println(arr[1])

	//指向数组的指针
	var arr1 *[5]int
	arr1 = &arr
	arr1[1] = 2
	fmt.Println(arr[1])

	//数组保存的是指针
	var arr2 [5]*int
	arr2[0] = &arr[0]
	arr2[1] = &arr[1]
	arr2[2] = &arr[2]

	fmt.Println(arr[1])
}
