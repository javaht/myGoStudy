package main

import "fmt"

func test01() {
	// 声明时没有指定数组元素的值, 默认为零值
	var arr [5]int
	fmt.Println(arr)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
}

func test02() {
	// 直接在声明时对数组进行初始化
	var arr1 = [5]int{15, 20, 25, 30, 35}
	fmt.Println(arr1)

	// 使用短声明
	arr2 := [5]int{15, 20, 25, 30, 35}
	fmt.Println(arr2)

	// 部分初始化, 未初始化的为零值
	arr3 := [5]int{15, 20} // [15 20 0 0 0]
	fmt.Println(arr3)

	// 可以通过指定索引，方便地对数组某几个元素赋值
	arr4 := [5]int{1: 100, 4: 200}
	fmt.Println(arr4) // [0 100 0 0 200]

	// 直接使用 ... 让编译器为我们计算该数组的长度
	arr5 := [...]int{15, 20, 25, 30, 35, 40}
	fmt.Println(arr5)
}

func test03() {
	// 特别注意数组的长度是类型的一部分，所以 [3]int 和 [5]int 是不同的类型
	arr1 := [3]int{15, 20, 25}
	arr2 := [5]int{15, 20, 25, 30, 35}
	fmt.Printf("type of arr1 is %T\n", arr1)
	fmt.Printf("type of arr2 is %T\n", arr2)
}

func test04() {
	// 定义多维数组  3行2列
	arr := [3][2]string{
		{"1", "Go语言极简一本通"},
		{"2", "Go语言微服务架构核心22讲"},
		{"3", "从0到Go语言微服务架构师"}}
	fmt.Println(arr) // [[15 20] [25 22] [25 22]]
}

func arrLength() {
	arr := [...]string{"Go语言极简一本通", "Go语言微服务架构核心22讲", "从0到Go语言微服务架构师"}
	fmt.Println("数组的长度是:", len(arr)) //数组的长度是: 3
}

func showArr() {
	arr := [...]string{"Go语言极简一本通", "Go语言微服务架构核心22讲", "从0到Go语言微服务架构师"}
	for index, value := range arr {
		fmt.Printf("arr[%d]=%s\n", index, value)
	}

	//_表示忽略索引,不需要索引值
	for _, value := range arr {
		fmt.Printf("value=%s\n", value)
	}
}

// Go 中的数组是值类型而不是引用类型。当数组赋值给一个新的变量时，该变量会得到一个原始数组的一个副本。如果对新变量进行更改，不会影响原始数组。
func arrByValue() {
	arr := [...]string{"Go语言极简一本通", "Go语言微服务架构核心22讲", "从0到Go语言微服务架构师"}
	copy := arr
	copy[0] = "Golang"
	fmt.Println(arr)
	fmt.Println(copy)
}

func main() {
	arrByValue()
}
