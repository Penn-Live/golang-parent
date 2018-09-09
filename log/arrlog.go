package smartlog

import "fmt"

//控制台打印log
func PrintArr(arr []string) {
	fmt.Println("length:", len(arr))
	fmt.Printf("%v", arr)
}

func LogArrtoFile(arr []string) {

}
