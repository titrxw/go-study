package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var test byte = 'A'
	fmt.Print(test) //通过ascll码表转换
	fmt.Print(string(test))

	var test1 byte = 65
	fmt.Print(test1)
	fmt.Print(string(test1))

	var num uint32 = 1010000000
	fmt.Print(string(num))

	//go 汉字占三个字节
	var name uint32 = '国'
	fmt.Print(name)
	fmt.Print(string(name))
	fmt.Print(unsafe.Sizeof(name))

	//go 汉字占三个字节，这里两个字节就可以表示是因为该字较简单，在unicode编码中国的排序靠前
	var name1 uint16 = '国'
	fmt.Print(name1)
	fmt.Print(string(name1))
	fmt.Print(unsafe.Sizeof(name1))

	var country string = "hello,中国"
	fmt.Println(len(country)) //5 + 1 + 6

	var str string = "test"
	var data []byte = []byte(str)
	fmt.Print(str)
	fmt.Print(data)

	var data1 [10]byte
	data1[0] = 'T'
	data1[1] = 'E'
	var str1 string = string(data1[:])
	fmt.Print(str1)
}
