package depchantable
// package main

import (
   "fmt"
    "data"
)


type DepChanTable struct{
	TID_RecTable map[int][] int //TID : Vector of TID dep
	//TID_ChanTable	map[int]chan data.Data //TID : Chan
} 

//for now static; will need taskSize later
func (table DepChanTable)SET_TID_RecTable(){
	var _0 = make([]int,3)
	_0[0] = 1
	_0[1] = 2
	_0[2] = 3

	var _1 = make([]int,2)
	_1[0] = 4
	_1[1] = 5

	var _3 = make([]int,1)
	_3[0] = 7

	var _4 = make([]int,1)
	_4[0] = 8

	var _5 = make([]int,1)
	_5[0] = 8

	var _6 = make([]int,1)	
	_6[0] = 2
	
	var _7 = make([]int,1)
	_7[0] = 8

	var _8 = make([]int,0)
	
	table.TID_RecTable[0] = _0
	table.TID_RecTable[1] = _1
	table.TID_RecTable[2] = _1
	table.TID_RecTable[3] = _1
	table.TID_RecTable[4] = _4
	table.TID_RecTable[5] = _5
	table.TID_RecTable[6] = _6
	table.TID_RecTable[7] = _7
	table.TID_RecTable[8] = _8
}

//for now static; will need taskSize later
// func (table DepChanTable)SET_TID_ChanTable(){
// 	var chan0_1 = make(chan data.Data)
// 	var chan0_2 = make(chan data.Data)
// 	var chan0_3 = make(chan data.Data)
// 	var chan1 = make(chan data.Data)
// 	var chan2 = make(chan data.Data)
// 	var chan3 = make(chan data.Data)
// 	var chan4 = make(chan data.Data)
// 	var chan5 = make(chan data.Data)
// 	var chan6 = make(chan data.Data)
// 	var chan7 = make(chan data.Data)

// 	table.TID_ChanTable[0] = chan0
// 	table.TID_ChanTable[1] = chan1
// 	table.TID_ChanTable[2] = chan2
// 	table.TID_ChanTable[3] = chan3
// 	table.TID_ChanTable[4] = chan4
// 	table.TID_ChanTable[5] = chan5
// 	table.TID_ChanTable[6] = chan6
// 	table.TID_ChanTable[7] = chan7

// }

func(table DepChanTable)PrintTIDTable(){
	for i, deps := range table.TID_RecTable{
		fmt.Println("TID ", i, "Dep(s): ",deps)
	}
}

// func(table DepChanTable)PrintChanTable(){
// 	for i, chanElem := range table.TID_ChanTable{
// 		fmt.Println("chan ", i, " ",chanElem)
// 	}
}
// func main(){
// 	TID_RecTable := make(map[int][]int)
// 	TID_ChanTable := make(map[int]chan data.Data)
// 	test := DepChanTable{TID_RecTable,TID_ChanTable}
// 	test.SET_TID_RecTable()
// 	fmt.Println("8th's first ",test.TID_RecTable[8][0])
// 	test.SET_TID_ChanTable()
// 	fmt.Println("Channel 5 ",test.TID_ChanTable[5])
// }