package util

func F(B uint32, C uint32, D uint32) uint32{
        //(B AND C) OR ((NOT B) AND D)
        return (B & C) | ((^B) & D)
}

func G(B uint32, C uint32, D uint32) uint32{
        //(B AND D) OR (C AND NOT D)
        return (D & B) | (C & (^D))
}

func H(B uint32, C uint32, D uint32) uint32{
        ///B XOR C XOR D
        return B ^ C ^ D
}

func I(B uint32, C uint32, D uint32) uint32{
        //C XOR (B OR (NOT D))
        return C ^ (B | (^D))
}

func J(B uint32, C uint32, D uint32) uint32{
	//(B AND C) OR (B AND D) OR (C AND D)
	return (B & C) | (B & D) | (C & D)
}
