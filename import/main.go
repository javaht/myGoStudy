package main //声明main包
//import "fmt" //导入fmt包 打印字符串需要

//导入多个包可以这么写
import (
	"fmt"
	"os"
)

// 声明main主函数 只能在main包中出现
func main() {
	fmt.Println("hello Go 这是我第一个Go程序")
	os.Getgid()
}

//func  函数名(参数列表)(返回值列表){
//  //执行体
// }
