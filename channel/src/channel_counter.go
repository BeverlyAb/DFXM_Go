package main

import (
    "fmt"
    "strings"
)

type Result struct {
    msg string
    seqNum  int
}

type Receiver struct{
	value string
	seqNum int
} 

func transmit(words []string, c chan Result) {
    defer close(c)

    seqNum := 0
    for _, word := range words {
        res := new(Result)
        res.msg = strings.ToUpper(word)
        res.seqNum = seqNum

        //seqNum++
        c <- *res       
    }
}

func updateValue(receivedVal string, senderSeqNum int, rec * Receiver){
	if rec.seqNum == senderSeqNum{
		rec.value = receivedVal
		rec.seqNum = rec.seqNum + 1
	} 
}

func main() {
    rec := new(Receiver)
    rec.value = "NULL"
    rec.seqNum = 0

    words := []string{"fire", "refire1", "refire2", "refire3", "refire4"}
    c := make(chan Result)
    go transmit(words, c)
    for res := range c {
        fmt.Print("SENT= ", res.msg, ", Updated Value= ") //, ",", res.seqNum
        updateValue(res.msg, res.seqNum,rec)
        fmt.Println(rec.value)
    }
}
/*

*/