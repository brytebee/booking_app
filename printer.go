package main

import "fmt"

// strings
var str = "Jon"
// List
var someList = []string{"Jon", "Snow", "jsnow@got.tv"}
// map
var someMap = make(map[string]string)
// list of maps
var listOfMaps = make([]map[string]uint, 0)
// list of string
var listOfStucts = make([]SomeStruct, 0)

// Struct
type SomeStruct struct {
	prop1 string
	prop2 uint
	prop3 int
	prop4 bool
}

var myStruct = SomeStruct {
	prop1: "",
	prop2: 3,
	prop3: -5,
	prop4: false,
}

func printer() {
	fmt.Printf("str = %v\n", str)
	fmt.Printf("someList = %v\n", someList)
	fmt.Printf("someMap = %v\n", someMap)
	fmt.Printf("ListOfMaps = %v\n",listOfMaps)
	fmt.Printf("myStruct = %v\n", myStruct)
	fmt.Printf("List of structs = %v\n", listOfStucts)
}