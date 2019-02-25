package main

import (
	"fmt"
	"github.com/GodlikePenguin/FunctionalGo/closures"
	"github.com/GodlikePenguin/FunctionalGo/slices"
)

func main() {
	//Slice example code
	fmt.Println("-----Slice examples-----")
	slices.IntExample()
	slices.StringExample()

	//Closure example code
	fmt.Println("-----Closure examples-----")
	closures.CurryExample()
	closures.ClosureExample()
}
