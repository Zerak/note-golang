package main

import "fmt"

const (
	loc0, bit0 uint32 = iota, 1 << iota   //loc0=0,bit0=1
	loc1, bit1                            //loc1=1,bit1=2
	loc2, bit2                            //loc2=2,bit2=4
)
//loc0 0  bit0 1  loc1 1  bit1 1  loc2 2  bit2 1
//loc0 0  bit0 1  loc1 1  bit1 2  loc2 2  bit2 4

const (
	h = iota    //h=0
	i = 100     //i=100
	j           //j=100
	k = iota    //k=3
)

func main() {
	fmt.Println("loc0", loc0," bit0", bit0, " loc1", loc1, " bit1", bit1, " loc2", loc2, " bit2", bit2)

	fmt.Println(h,i,j,k)
}
