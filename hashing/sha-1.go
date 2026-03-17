package main

import ("fmt"; "hashing/util"; "math/bits"; "os")

func operation(A uint32, B uint32, C uint32, D uint32, E uint32, word uint32, round int) uint32{
	var functionResult uint32 = 0
	var f, k uint32 = 0, 0
	var quotient int = round / 20
	switch quotient{
	case 0:
		f = util.F(B,C,D)
		k = 0x5A827999
	case 1:
		f = util.I(B,C,D)
		k = 0x6ED9EBA1
	case 2:
		f = util.J(B,C,D)
		k = 0x8F1BBCDC
	case 3:
		f = util.I(B,C,D)
		k = 0xCA62C1D6
	}
	functionResult = uint32(word+functionResult)
	functionResult = uint32(k+functionResult)
	functionResult = uint32(E+functionResult)
	functionResult = uint32(f+functionResult)
	functionResult = uint32(bits.RotateLeft32(A,5)+functionResult)
	return functionResult
}


func main(){
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


        data = util.Padding(data, false)
	//Everything for SHA-1 is Big Endian

	/*
        The starting values are the following turned into little endian which I have done manually
        0x01234567
        0x89abcdef
        0xfedcba98
        0x76543210
	0xF0E1D2C3
        */
        var initialA uint32 = 0x67452301
        var initialB uint32 = 0xefcdab89
        var initialC uint32 = 0x98badcfe
        var initialD uint32 = 0x10325476
	var initialE uint32 = 0xC3D2E1F0

	var A uint32 = initialA
	var B uint32 = initialB
	var C uint32 = initialC
	var D uint32 = initialD
	var E uint32 = initialE

        var new_A, new_B, new_C, new_D, new_E uint32 = 0, 0, 0, 0, 0

	for i:=0;i<len(data)/64;i++{
		M:=data[64*i:64*(i+1)] //512-bit block
		words := make([]uint32,80)
		for j:=0;j<16;j++{
			words[j] = util.BytesToInt32(M[4*j:4*(j+1)])
		}
		for j:=16;j<80;j++{
			words[j] = bits.RotateLeft32((words[j-3] ^ words[j-8] ^ words[j-14] ^ words[j-16]), 1)
		}

		var blockInitialA uint32 = A
                var blockInitialB uint32 = B
                var blockInitialC uint32 = C
                var blockInitialD uint32 = D
		var blockInitialE uint32 = E

		for round:=0;round<80;round++{
			word := words[round]
			new_A = operation(A, B, C, D, E, word, round)
			new_B = A
			new_C = bits.RotateLeft32(B,30)
			new_D = C
			new_E = D

			A = new_A
                        B = new_B
                        C = new_C
                        D = new_D
			E = new_E
		}
		//Feed forward step
                new_A = uint32(A+blockInitialA)
                new_B = uint32(B+blockInitialB)
                new_C = uint32(C+blockInitialC)
                new_D = uint32(D+blockInitialD)
		new_E = uint32(E+blockInitialE)
                A = new_A
                B = new_B
                C = new_C
                D = new_D
		E = new_E
	}
	fmt.Printf("Result:\n%08x%08x%08x%08x%08x\n", A, B, C, D, E)
}
