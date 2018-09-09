package arrlog

import "fmt"

//控制台打印log
func LogArr(arr []string) {
	fmt.Println("length:", len(arr))
	fmt.Printf("%v", arr)
}

func FileLogArr(arr []string) {

}
