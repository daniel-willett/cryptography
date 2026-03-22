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
	
	keyData, err := os.ReadFile("permutation-cipher-key")
	if err != nil{
        	panic(err)
        }
	keyDataString := string(keyData)
	
	blockLength := common.Split(keyDataString,"\n")[0]
	keyPermutations := common.Split(keyDataString,"\n")[1]

	/* To do:
	We need to verify that the `blockLength` matches as a factor of the number of characters in the plaintext
	I think a sufficiently true permutation cipher should permute the special characters as well as text but I don't think most people to be in the spirit of what the permutation cipher is supposed to be from an intelectual standpoint.
	So to only permute text I think we'd need another variable with the text minus all the special characters.
	--> then apply the cipher to it
	--> then go back through the old text using two pointers/indexes to replace old cases of text with the new ones.
	*/
}
