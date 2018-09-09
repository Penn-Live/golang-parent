package file

import (
	"io/ioutil"
	"log"
	"strings"
)

/**
  根据某种分隔符来读取文件，返回字符数组
*/
func ReadFileWithDelimiter(fileName string, delimiter string) []string {

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("file:", fileName, " does not exit.")
	}
	content := strings.TrimSpace(string(bytes))
	if content == "" {
		println("file is empty.")
		return nil
	}
	contentArr := strings.Split(content, delimiter)
	println("return arr length :", len(contentArr))
	return contentArr
}
