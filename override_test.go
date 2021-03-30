package typoverride_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"typoverride"
)

type ExampleT struct {
	InputType           string
	ValueType           string
	ValueUnit           string
	SelectList          []string
	SelectListInterface string
}
var _ = typoverride.Do(ExampleT{})

func Test1(t *testing.T){

	var a = ExampleT{
		InputType:           "1",
		ValueType:           "2",
		ValueUnit:           "3",
		SelectList:          []string{"4.1", "4.2"},
		SelectListInterface: "5",
	}
	data, _ := json.Marshal(a)
	fmt.Println(string(data)) // 请看输出的 json key 的格式
}
