package main

import ("fmt"; "hashing/util"; "math/bits"; "os")

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
		words := make([][]byte,16)
		for j:=[0];j<16;j++{
			words[j] = M[4*j:4*(j+1)]
		}
		for j:=16;j<80;j++{
			words[j] = bits.RotateLeft32((words[j-3] ^ words[j-8] ^ words[j-14] ^ words[j-16]), 1)
		}


	}
}
