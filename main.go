package main

import "fmt"

type Test struct {
	value int
}

func main() {

	test := Test{
		value: 1,
	}
	fmt.Printf("Main Pointer %p\n", &test)

	OnlyRead(test)
	fmt.Println("Main after OnlyRead: ", test)
	Update(&test)
	fmt.Printf("Pointer after OnlyRead %p\n", &test)
	fmt.Println("Main after Update: ", test)
	fmt.Printf("Pointer after Update %p\n", &test)
}

func OnlyRead(testValue Test) {
	fmt.Printf("OnlyRead Pointer %p\n", &testValue)

	testValue.value = 7
	fmt.Println("OnlyRead: ", testValue)
}

func Update(testValue *Test) {
	fmt.Printf("Update Pointer %p\n", testValue)
	testValue.value = 3
	fmt.Println("Update: ", testValue)
}
