package main


func SetBit(n int64, i uint, bit bool)(int64){
	if bit{
		return n | (1 << i)
	} else{
		return  n &^ (1 << i)
	}
}
