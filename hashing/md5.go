package main

import ("fmt")

func padding(data []byte) []byte{
	//Each []byte is 256 bits as []byte takes values 0-255
	bitLength := len(data) * 256
}


func main(){
	input := "Hello World"
	data := []byte(input)
	data = padding(data)
}
