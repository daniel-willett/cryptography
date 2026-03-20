package main

import ("fmt"; "os"; "github.com/daniel-willett/common-go-methods")

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
	
	keyData, err := os.ReadFile("substitution-cipher-key")
	if err != nil{
        	panic(err)
        }
	keyDataString := string(keyData)
	
	var result string
	for i:=0; i<len(data); i++{
		val := data[i]
		if (val>=97 && val<=122){
			val -= 32
		}
		char := string(val)
		pos := common.GetIndexOf(keyDataString, char)
		if len(pos)==0{
			result += char
		} else {
			result += string(keyDataString[pos[0]+26+1]) //Go accross 16 characters in to the what the char gets encrypted to, plus 1 for the '\n' character
		}
	}
	fmt.Println(result)
	
}
