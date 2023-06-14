package main

import "fmt"

// misc code snippets - for struct declaration
// refer to subsection in chapter 2 about structs

func main() {
	testVar := TestStruct{Sample: "aa", SampleInt: 1}
	fmt.Println(testVar)

	testVar2 := TestStruct{"aa", 1}
	fmt.Println(testVar2)

	testVar3 := TestStruct2{"aa", 1}
	fmt.Println(testVar3)
	testVar3.ChangeSample("bb")
	fmt.Println(testVar3)

	testVar4 := TestStruct3{"aa", 1}
	fmt.Println(testVar4)
	testVar4.ChangeSample("bb")
	fmt.Println(testVar4)

}

type TestStruct struct {
	Sample    string
	SampleInt int
}

type TestStruct2 struct {
	Sample    string
	SampleInt int
}

func (t TestStruct2) ChangeSample(s string) {
	t.Sample = s
	fmt.Println(t)
}

type TestStruct3 struct {
	Sample    string
	SampleInt int
}

func (t *TestStruct3) ChangeSample(s string) {
	t.Sample = s
	fmt.Println(t)
}

type TestStruct5 struct {
	Sample    string `json:"sample"`
	SampleInt int    `json:"sample_int"`
}
