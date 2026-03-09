package main

import ("fmt"; "math")

func F(B uint32, C uint32, D uint32) uint32{
	//(B AND C) OR ((NOT B) AND D)
	//NOT B is the same as 11111111 XOR B
	return (B & C) | ((11111111 ^ B) & D)
}

func G(B uint32, C uint32, D uint32) uint32{
	//(B AND C) OR (C AND NOT D)
        //NOT D is the same as 11111111 XOR D
        return (B & C) | (C & (11111111^D))
}

func H(B uint32, C uint32, D uint32) uint32{
	///B XOR C XOR D
	return B ^ C ^ D
}

func I(B uint32, C uint32, D uint32) uint32{
	//C XOR (B OR (NOT D))
	//NOT D is the same as 11111111 XOR D
	return C ^ (B | (11111111^D))
}


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

func bytesToInt32(word []byte) uint32{
	//[01 23 45 67] into a single value (presumably in decimal as the []byte is in decimal)
	//
	var first, second, third, fourth uint32 = 0, 0, 0, 0
	first = uint32(word[0]) * uint32(math.Pow(2,8*3))
        second = uint32(word[1]) * uint32(math.Pow(2,8*2))
        third = uint32(word[2]) * uint32(math.Pow(2,8*1))
        fourth = uint32(word[3]) * uint32(math.Pow(2,8*0))
	return first + second + third + fourth
}

func operation(A uint32, B uint32, C uint32, D uint32, word []byte, K uint32, round int) uint32{
	var functionResult uint32 = 0
	switch round{
	case 0:
		functionResult = F(B,C,D)
	case 1:
		functionResult = G(B,C,D)
	case 2:
		functionResult = H(B,C,D)
	case 3:
		functionResult = I(B,C,D)
	}
	functionResult = uint32(A+functionResult) //A+F(B,C,D) mod 2^32
	converted := bytesToInt32(word)
	functionResult = uint32(converted+functionResult) //A+F(B,C,D)+M_i mod 2^32
	functionResult = uint32(K+functionResult) //A+F(B,C,D)+M_i+K_i mod 2^32
}

func KFormula(round int,operationNumber int)uint32{
	var i float64 = float64((16*round)+operationNumber)
	sineValue := math.Sin(angle)
	value := math.Abs(sineValue) * math.Pow(2,32)
	quotient := int(math.Trunc(value))
	return uint32(quotient)
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

	var A uint32 = 0x01234567
        var B uint32 = 0x89abcdef
        var C uint32 = 0xfedcba98
        var D uint32 = 0x76543210

	var new_A, new_B, new_C, new_D uint32 = 0, 0, 0, 0

	rounds := [][]int{
		{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15},
		{1,6,11,0,5,10,15,4,9,14,3,8,13,2,7,12},
		{5,8,11,14,1,4,7,10,13,0,3,6,9,12,15,2},
		{0,7,14,5,12,3,10,1,8,15,6,13,4,11,2,9},
	}

	for i:=0;i<len(data)/64;i++{
		M:=data[i:i+64] //512-bit blocks
		words := make([][]byte,16)
		for j:=0;j<16;j++{
			words[j] = M[4*j:4*(j+1)] //In each 512-bit Block we have 16 32-bit words (M_0 to M_15)
		}
		fmt.Println(words)
		
		for round:=0;round<4;round++{
			for operationNumber:=0;operationNumber<16;operationNumber++{
				K = KFormula(round,operationNumber)
				word := words[rounds[round][operationNumber]] //This is a [4]byte
				new_A = D
				new_B = operation(A,B,C,D,word,K,round)
				new_C = B
				new_D = C

				A = new_A
				B = new_B
				C = new_C
				D = new_D
			}
		}
	}
}
