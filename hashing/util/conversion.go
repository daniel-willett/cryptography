package util

import "math"

func BytesToInt32(word []byte) uint32{
        var total uint32 = 0
        for i:=0; i<4; i++{
                total += uint32(word[i]) * uint32(math.Pow(2,float64(8*(3-i))))
        }
        return total
}
