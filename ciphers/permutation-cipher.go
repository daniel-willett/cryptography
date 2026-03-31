package main

import ("fmt"; "os"; "github.com/daniel-willett/common-go-methods")

func strip(arr []byte) []byte{
	var result []byte
	for _, val := range arr{
		if (val>=65 && val<=90){
			result = append(result, val)
		}
	}
	return result
}

func encrypted(textBlock []byte, keyDataString string, blockLength int) []byte {
	var result []byte
	for i:=0; i<blockLength; i++{
		//keyDataString[i] is a `byte` so we subtract 48 to get the numerical value. We -1 because the key written by a human is indexed 1 to n rather than 0 to n-1
		result = append(result, textBlock[keyDataString[2*i]-1-48])
	}
	return result
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
	
	data = []byte(common.Upper(string(data)))

	keyData, err := os.ReadFile("permutation-cipher-key")
	if err != nil{
        	panic(err)
        }
	keyDataString := string(keyData)
	
	originalText := string(data)
	justText := strip(data)

	fmt.Println(string(justText))

	numberOfChars := len(justText)
	blockLength := len(common.Split(keyDataString," "))

	if numberOfChars%blockLength!=0{
		fmt.Println("Error: Text length doesn't match with factor of key block")
		os.Exit(1)
	}

	numberOfBlocks := numberOfChars/blockLength
	var result []byte
	for i:=0;i<numberOfBlocks;i++{
		textBlock := justText[blockLength*i:blockLength*(i+1)]
		result = append(result, encrypted(textBlock, keyDataString, blockLength)...)
	}

	fmt.Println(string(result))

	var ciphertext string
	var compressedPointer int = 0
	for i:=0; i<len(originalText); i++{
		val:=originalText[i]
		if (val>=65 && val<=90){
			ciphertext += string(result[compressedPointer])
			compressedPointer += 1
		} else {
			ciphertext += string(val)
		}
	}

	fmt.Println(ciphertext)
}
