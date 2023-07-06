package main

import "fmt"

func greetThem(name string) {
	fmt.Printf("Hey %v \n", name)
}

func anyMessage() string {
	return "How are you"
}

func addThree(num1, num2, num3 uint) uint {
	return num1 + num2 + num3
}

func anyNumber() uint{
	return 7
}

func anyTwoNum() (uint, uint){
	return 3, 4
}

func main() {
	greetThem("Musa")
	fmt.Println(anyMessage())
	num_a, num_b := anyTwoNum();
	fmt.Println(addThree(num_a, num_b, anyNumber()))
}