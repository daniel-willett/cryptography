package util

import "math"

func BytesToInt32(word []byte, littleEndian bool) uint32{
        var total uint32 = 0
        for i:=0; i<4; i++{
		switch littleEndian{
		case true:
			total += uint32(word[i]) * uint32(math.Pow(2,float64(8*(3-i))))
		case false:
			total += uint32(word[i]) * uint32(math.Pow(2,float64(8*i)))
		}
        }
        return total
}
