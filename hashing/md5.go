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

	A := 0x01234567
        B := 0x89abcdef
        C := 0xfedcba98
        D := 0x76543210

	rounds := [][]int{
		{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15},
		{1,6,11,0,5,10,15,4,9,14,3,8,13,2,7,12},
		{5,8,11,14,1,4,7,10,13,0,3,6,9,12,15,2},
		{0,7,14,5,12,3,10,1,8,15,6,13,4,11,2,9},
	}

	for i:=0;i<len(data)/64;i++{
		M:=data[i:i+64]
		words := make([][]byte,16)
		for j:=0;j<16;j++{
			words[j] = M[4*j:4*(j+1)]
		}
		fmt.Println(words)
	}
}
