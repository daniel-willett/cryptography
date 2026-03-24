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

func encrypted(textBlock []byte, keyDataString string) []byte {

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
	
	//originalText := string(data)
	justText := strip(data)

	numberOfChars := len(justText)
	blockLength := len(common.Split(keyDataString," "))

	if numberOfChars%blockLength==0{
		fmt.Println("Error")
		os.Exit(1)
	}

	numberOfBlocks := numberOfChars/blockLength
	var result []byte
	for i:=0;i<numberOfBlocks;i++{
		textBlock := justText[blockLength*i:blockLength*(i+1)]
		result = append(result, encrypted(textBlock, keyDataString)...)
	}
}
