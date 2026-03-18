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
	/*
	fmt.Printf("ROTL5(A) is %08x\n", bits.RotateLeft32(A,5))
	fmt.Printf("f is %08x\n", f)
	fmt.Printf("E is %08x\n", E)
	fmt.Printf("W[0] is %08x\n", word)
	fmt.Printf("K is %08x\n", k)
	*/
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
        var h0 uint32 = 0x67452301
        var h1 uint32 = 0xefcdab89
        var h2 uint32 = 0x98badcfe
        var h3 uint32 = 0x10325476
	var h4 uint32 = 0xC3D2E1F0

	var A, B, C, D, E uint32 = 0

	for i:=0;i<len(data)/64;i++{
		M:=data[64*i:64*(i+1)] //512-bit block
		words := make([]uint32,80)
		for j:=0;j<16;j++{
			words[j] = util.BytesToInt32(M[4*j:4*(j+1)]) 
		}
		fmt.Printf("%08x\n",words)
		for j:=16;j<80;j++{
			words[j] = bits.RotateLeft32((words[j-3] ^ words[j-8] ^ words[j-14] ^ words[j-16]), 1)
		}

		A = h0
		B = h1
		C = h2
		D = h3
		E = h4

		for round:=0;round<80;round++{
			word := words[round]
			temp := operation(A, B, C, D, E, word, round)
			E = D
			D = C
			C = bits.RotateLeft32(B, 30)
			B = A
			A = temp
			/*
			fmt.Printf("E is %08x\n", E)
			fmt.Printf("D is %08x\n", D)
			fmt.Printf("C is %08x\n", C)
			fmt.Printf("B is %08x\n", B)
			fmt.Printf("A is %08x\n", A)
			*/
		}
		h0 = uint32(A+h0)
		h1 = uint32(B+h1)
		h2 = uint32(C+h2)
		h3 = uint32(D+h3)
		h4 = uint32(E+h4)
	}
	fmt.Printf("Result:\n%08x%08x%08x%08x%08x\n", h0, h1, h2, h3, h4)
}
