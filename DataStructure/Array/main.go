package main

import "fmt"

type myArray struct {
	size     int
	capacity int
	array    mineArray
}

type mineArray interface {
	toString() string
	setValue(key int)
	getValue(key int) int
}

type baseArray struct {

}

func (b *baseArray)toString ()string  {
	return "toStringTest"
}
func  (b *baseArray)setValue(key int)  {

}
func  (b *baseArray)getValue(key int) int {
	return 1
}

func (f *myArray) add(value int) {
	f.size ++
	if f.size == f.capacity {
		f.capacity *= 2
	}
}

func (f *myArray) getSize() int {
	return f.size
}

func (f *myArray) toString() string {
	str := " "
	return str
}



func createArray() myArray {
	return myArray{0, 10,new(baseArray),}

}

func main() {
	array := createArray()
	array.add(1)
	array.add(2)
	fmt.Println(array.toString())

}
