//package depchantable
package main

import (
    "fmt"
    "data"
)


type DepChanTable struct{
	TID_DEPTable map[int][] int //TID : Vector of TID dep
	TID_ChanTable	map[int]chan data.Data //TID : Chan
} 

//for now static; will need taskSize later
func (table DepChanTable)SET_TID_DEPTable(){
	var _0 = make([]int,0)
	var _123 = make([]int,1)
	_123[0] = 0

	var _4 = make([]int,1)
	_4[0] = 1

	var _5 = make([]int,2)
	_5[0] = 1
	_5[1] = 2

	var _6 = make([]int,1)	
	_6[0] = 2
	
	var _7 = make([]int,1)
	_7[0] = 3

	var _8 = make([]int,4)		
	_8[0] = 4
	_8[1] = 5
	_8[2] = 6
	_8[3] = 7
	
	table.TID_DEPTable[0] = _0
	table.TID_DEPTable[1] = _123
	table.TID_DEPTable[2] = _123
	table.TID_DEPTable[3] = _123
	table.TID_DEPTable[4] = _4
	table.TID_DEPTable[5] = _5
	table.TID_DEPTable[6] = _6
	table.TID_DEPTable[7] = _7
	table.TID_DEPTable[8] = _8
}

//for now static; will need taskSize later
// func (d DepChanTable)TID_ChanTable(){

// }

func main(){
	TID_DEPTable := make(map[int][]int)
	TID_ChanTable := make(map[int]chan data.Data)
	test := DepChanTable{TID_DEPTable,TID_ChanTable}
	test.SET_TID_DEPTable()
	fmt.Println("8th's first ",test.TID_DEPTable[8][0])

}