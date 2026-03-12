package main

import ("fmt"; "math"; "math/bits"; "time"; "os")

func F(B uint32, C uint32, D uint32) uint32{
	//(B AND C) OR ((NOT B) AND D)
	return (B & C) | ((^B) & D)
}

func G(B uint32, C uint32, D uint32) uint32{
	//(B AND D) OR (C AND NOT D)
        return (D & B) | (C & (^D))
}

func H(B uint32, C uint32, D uint32) uint32{
	///B XOR C XOR D
	return B ^ C ^ D
}

func I(B uint32, C uint32, D uint32) uint32{
	//C XOR (B OR (NOT D))
	return C ^ (B | (^D))
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
	lengthPadding := make([]byte, 64/8) //64bit but each byte is 8bits so we need 64/8=8 many bytes
	//[0     0     0     0     0     0     0     0]
	//2^8^7 2^8^6 2^8^5 2^8^4 2^8^3 2^8^2 2^8^1 2^8^0
	for i:=7; i>=0; i--{
		var temp int = lengthPaddingValue / int(math.Pow(2,float64(8*i)))
		lengthPadding[7-i] = byte(temp)
		lengthPaddingValue = lengthPaddingValue - (temp*int(math.Pow(2,float64(8*i))))
	}
	lengthPadding = littleEndian(lengthPadding)
	data = append(data, lengthPadding...)
	return data
}

func bytesToInt32(word []byte) uint32{
	var total uint32 = 0
	for i:=0; i<4; i++{
		total += uint32(word[i]) * uint32(math.Pow(2,float64(8*(3-i))))
	}
	return total
}

func operation(A uint32, B uint32, C uint32, D uint32, word []byte, K uint32, round int, S int) uint32{
	var functionResult uint32 = 0
	converted := bytesToInt32(word)
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
	functionResult = uint32(A+functionResult) 		// A+F(B,C,D) mod 2^32
	functionResult = uint32(converted+functionResult) 	// A+F(B,C,D)+M_i mod 2^32
	functionResult = uint32(K+functionResult) 		// A+F(B,C,D)+M_i+K_i mod 2^32
	functionResult = bits.RotateLeft32(functionResult, S) 	// <<<S
	functionResult = uint32(B+functionResult) 		// ((A+F(B,C,D)+M_i+K_i mod 2^32)<<<S)+B mod 2^32
	return functionResult
}

func KFormula(round int,operationNumber int)uint32{
	var i float64 = float64((16*round)+operationNumber)
	sineValue := math.Sin(i+1)
	value := math.Abs(sineValue) * math.Pow(2,32)
	quotient := int(math.Trunc(value))
	return uint32(quotient)
}

func littleEndian(words []byte) []byte{
	n := len(words)
	newOne := []byte{}
	for i:=0; i<n; i++{
		newOne = append(newOne, words[n-1-i])
	}
	return newOne
}

func fromLittleEndian(x uint32)uint32{
	//[A, B, C, D]
	//To clear the values left and right of each value we push it to the leftmost, <<8*i, and then right most, >>24, positions
	//Then move it to the position we want it to go, <<8*i
	var total uint32 = 0
	for i:=0; i<4; i++{
		total += ((x<<(8*i)>>24)<<(8*i))
	}
	return total
}

func main(){
	
	start := time.Now()

	args := os.Args
	var data []byte
	var input string = "Hello World"
	var err error
	if len(args) != 1{
		data, err = os.ReadFile(args[1])
        	if err != nil{
                	panic(err)
        	}
	} else {
		data = []byte(input)
	}
	data = padding(data)

	/*
	The starting values are the following turned into little endian which I have done manually
	0x01234567
        0x89abcdef
        0xfedcba98
        0x76543210
	*/
	var initialA uint32 = 0x67452301
	var initialB uint32 = 0xefcdab89
	var initialC uint32 = 0x98badcfe
	var initialD uint32 = 0x10325476
	
	var A uint32 = initialA
	var B uint32 = initialB
	var C uint32 = initialC
	var D uint32 = initialD

	var new_A, new_B, new_C, new_D uint32 = 0, 0, 0, 0

	rounds := [][]int{
		{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15},
		{1,6,11,0,5,10,15,4,9,14,3,8,13,2,7,12},
		{5,8,11,14,1,4,7,10,13,0,3,6,9,12,15,2},
		{0,7,14,5,12,3,10,1,8,15,6,13,4,11,2,9},
	}

	shifts := [][]int{
		{7,12,17,22,7,12,17,22,7,12,17,22,7,12,17,22},
		{5,9,14,20,5,9,14,20,5,9,14,20,5,9,14,20},
		{4,11,16,23,4,11,16,23,4,11,16,23,4,11,16,23},
		{6,10,15,21,6,10,15,21,6,10,15,21,6,10,15,21},
	}

	for i:=0;i<len(data)/64;i++{
		M:=data[64*i:64*(i+1)] //512-bit blocks
		words := make([][]byte,16)
		//In each 512-bit Block, M, we have 16 32-bit words (M_0 to M_15)
		for j:=0;j<16;j++{
			words[j] = littleEndian(M[4*j:4*(j+1)])
		}

		//The initial vectors are the values from the previous block (in the case this is the first block then we use the original values declared before)
		var blockInitialA uint32 = A
		var blockInitialB uint32 = B
		var blockInitialC uint32 = C
		var blockInitialD uint32 = D

		for round:=0;round<4;round++{
			for operationNumber:=0;operationNumber<16;operationNumber++{
				K := KFormula(round,operationNumber)
				S := shifts[round][operationNumber]
				word := words[rounds[round][operationNumber]] //This is a [4]byte
				new_A = D
				new_B = operation(A,B,C,D,word,K,round,S)
				new_C = B
				new_D = C

				A = new_A
				B = new_B
				C = new_C
				D = new_D
			}
		}
		//Feed forward step
		new_A = uint32(A+blockInitialA)
		new_B = uint32(B+blockInitialB)
		new_C = uint32(C+blockInitialC)
		new_D = uint32(D+blockInitialD)
		A = new_A
		B = new_B
		C = new_C
		D = new_D
	}
	end := time.Since(start)
	fmt.Printf("Time taken: %s\n", end)
	fmt.Printf("Result:\n%08x%08x%08x%08x\n", fromLittleEndian(A),fromLittleEndian(B),fromLittleEndian(C),fromLittleEndian(D))
}
