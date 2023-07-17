package main

import "fmt"

func main() {

	//var 关键字
	//方法一： 声明一个变量,默认值是0 类型是int
	//var a int
	//fmt.Println("a= ", a)
	//fmt.Printf("a的类型: %T\n", a)

	//方法二：声明一个变量，并且初始化一个值
	//var b int = 100
	//fmt.Printf("b=%d,b的类型是：%T", b, b)

	//方法三：初始化的时候，去掉数据类型。go语言通过值自动匹配类型
	//var c = 100
	//fmt.Printf("c=%d,c的类型是:%T", c, c)

	//var c = "哈哈哈哈哈哈"
	//fmt.Printf("c=%d,c的类型是:%T", c, c)

	//方法四：短声明 :=
	//e := 100
	//fmt.Printf("e=%d,c的类型是:%T", e, e)

	//多个变量声明方式：
	//var xx,yy int =100,200
	//var kk,wx =300,"qitaleixing"
	//var (
	//   nn int =100
	//   wechat = "qitaleixing"
	//)

	//常量定义使用const
	//const length = 10
	//fmt.Println("length=", length)

	//const枚举使用
	//const (
	//	BEIJING  = 0
	//	SHANGHAI = 1
	//	SHENZHEN = 2
	//)
	//fmt.Println("BEIJING=", BEIJING)
	//fmt.Println("SHANGHAI=", SHANGHAI)
	//fmt.Println("SHANGHAI=", SHANGHAI)

	//IOTA常量计数器
	const (
		BEIJING = 10 * iota
		SHANGHAI
		SHENZHEN
	)

	fmt.Println("BEIJING=", BEIJING) //代表在枚举中的从0开始 默认是每次递增1   BEIJING =0  SHANGHAI=1    也可以递增10   BEIJING =0  SHANGHAI=10
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("SHANGHAI=", SHENZHEN)

}
