package file_test

import (
	"../file"
	"../log"
	"testing"
)

func TestReadFileWithDelimiter(t *testing.T) {
	delimiter := file.ReadFileWithDelimiter("/Users/tangzhongping/Desktop/ids/clean_id.dat", ",")
	smartlog.PrintArr(delimiter)
}
