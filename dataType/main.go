package main

import (
	"fmt"
	"math"
	"unsafe"
)

// 有符号整型

func Integer() {
	var num8 int8 = 127
	var num16 int16 = 32767
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	var num int = math.MaxInt
	//unsafe.Sizeof(变量)会返回该变量类型所占用的字节数
	fmt.Printf("num8的类型是：%T,num8占用的字节数是%d,num8值是 %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是：%T,num16占用的字节数是%d,num16值是 %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是：%T,num32占用的字节数是%d,num32值是 %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是：%T,num64占用的字节数是%d,num64值是 %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是：%T,num占用的字节数是%d,num值是 %d\n", num, unsafe.Sizeof(num), num)
}

// 无符号整型
func unsignedInteger() {
	var num8 uint8 = 128
	var num16 uint16 = 32768
	var num32 uint32 = math.MaxUint32
	var num64 uint64 = math.MaxUint64
	var num uint = math.MaxUint
	fmt.Printf("num8的类型是 %T, num8占用的字节数 %d, num8是 %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是 %T, num16占用的字节数 %d, num16是 %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是 %T, num32占用的字节数 %d, num32是 %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是 %T, num64占用的字节数 %d, num64是 %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是 %T, num占用的字节数 %d, num是 %d\n", num, unsafe.Sizeof(num), num)
}

func showFloat() {
	var num1 float32 = math.MaxFloat32
	var num2 float64 = math.MaxFloat64
	fmt.Printf("num1的类型是%T,num1是%g\n", num1, num1) //小数点后六位
	fmt.Printf("num2的类型是%T,num1是%g\n", num2, num2) //小数点后位
}

func main() {
	Integer()
	println("==================================")
	unsignedInteger()
	println("==================================")
	showFloat()
}
