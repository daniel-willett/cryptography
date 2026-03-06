package main

import ("fmt"; "math")

func padding(data []byte) []byte{
	//Each []byte is 8 bits as []byte takes values 0-255
	bitLength := len(data) * 8
	var quotient int = (bitLength + 64) / 512 //We 'reserve' 64 at the end anyway and we need the total of the 0 padding & length padding to bring to a multiple of 512
	paddingNeeded := ((quotient+1)*512) - bitLength
	numberOfZeros := paddingNeeded - 64
	//First "chunk" (8bits) is taken by binary 10000000 which is 128 in decimal
	data = append(data, 128)
	for i:=1; i<=(numberOfZeros/8)-1; i++{ //We need all the "chunks" except the first one which has just been done
		data = append(data, 0)
	}
	lengthPaddingValue := bitLength%int(math.Pow(2,64))
	lengthPadding := make([]byte, 64/8) //64bit but each "chunk" is 8bits so we need 64/8=8 "chunks"
	//[0     0     0     0     0     0     0     0]
	//2^8^7 2^8^6 2^8^5 2^8^4 2^8^3 2^8^2 2^8^1 2^8^0
	for i:=7; i>=0; i--{
		var temp int = lengthPaddingValue / int(math.Pow(2,float64(8*i)))
		lengthPadding[7-i] = byte(temp)
		lengthPaddingValue = lengthPaddingValue - (temp*int(math.Pow(2,float64(8*i))))
	}
	data = append(data, lengthPadding...)
	return data
}


func main(){
	/*
	input := "Hello World"
	data := []byte(input)
	data = padding(data)
	fmt.Println(data)
	*/

	input := "They are deterministic"
	data := []byte(input)
	data = padding(data)
	fmt.Println(data)
}
