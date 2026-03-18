package util

func LittleEndianByteSlice(words []byte) []byte{
        n := len(words)
        newOne := []byte{}
        for i:=0; i<n; i++{
                newOne = append(newOne, words[n-1-i])
        }
        return newOne
}

func LittleEndianUInt32(x uint32) uint32{
        //[A, B, C, D]
        //To clear the values left and right of each value we push it to the leftmost, <<8*i, and then right most, >>24, positions
        //Then move it to the position we want it to go, <<8*i
        var total uint32 = 0
        for i:=0; i<4; i++{
                total += ((x<<(8*i)>>24)<<(8*i))
        }
	return total
}
