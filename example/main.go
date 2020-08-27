package main

import (
	"fmt"
	"github.com/jadefox10200/trivium"
)

func main() {

	t := trivium.NewTriv()

	//key must be 80 bits
	key := [10]uint8{120,120,120,120,120,120,120,120,120,120}

	t.LoadKey(key)

	//iv must be 80 bits
	iv := [10]uint8{1,165,120,120,120,120,120,120,120,120}
	err := t.Loadiv(iv)
	if err != nil {fmt.Println(err.Error()); return }

	//do first 1152 cycles to warm up.
	t.Init()

	//start outputting encryption bits:

	for i := 0; i < 64; i++ {
		fmt.Printf("%v", t.Clock())
	}

	return 
}