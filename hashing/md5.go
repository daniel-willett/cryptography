package main

import ("fmt"; "math"; "math/bits"; "time"; "os"; "hashing/util")

func operation(A uint32, B uint32, C uint32, D uint32, word []byte, K uint32, round int, S int) uint32{
	var functionResult uint32 = 0
	converted := util.BytesToInt32(word)
	switch round{
	case 0:
		functionResult = util.F(B,C,D)
	case 1:
		functionResult = util.G(B,C,D)
	case 2:
		functionResult = util.H(B,C,D)
	case 3:
		functionResult = util.I(B,C,D)
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


	data = util.Padding(data, true)

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
			words[j] = util.LittleEndianByteSlice(M[4*j:4*(j+1)])
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
	fmt.Printf("Result:\n%08x%08x%08x%08x\n", util.LittleEndianUInt32(A),util.LittleEndianUInt32(B),util.LittleEndianUInt32(C),util.LittleEndianUInt32(D))
	end := time.Since(start)
	fmt.Printf("Time taken: %s\n", end)
}
