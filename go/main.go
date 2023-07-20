package main

import (
	"fmt"
	"time"

	"github.com/HungTP-Play/ex/go/ex"
)

func main() {
	input := "adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc"
	startTime := time.Now().UnixMilli()
	count := ex.Ex1_2(input, "bcddceeeebecbc")
	endTime := time.Now().UnixMilli()
	fmt.Printf("time: %d\n", endTime-startTime)
	fmt.Printf("count: %d\n", count)
}
